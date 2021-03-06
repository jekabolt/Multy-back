/*
Copyright 2019 Idealnaya rabota LLC
Licensed under Multy.io license.
See LICENSE for details
*/
package btc

import (
	"context"
	"io"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/Appscrunch/Multy-back/currencies"
	pb "github.com/Appscrunch/Multy-back/node-streamer/btc"
	"github.com/Appscrunch/Multy-back/store"
	nsq "github.com/bitly/go-nsq"
	"gopkg.in/mgo.v2/bson"
)

func setGRPCHandlers(cli pb.NodeCommuunicationsClient, nsqProducer *nsq.Producer, networtkID int, wa chan pb.WatchAddress) {

	// initial fill mempool respectively network id
	go func() {
		stream, err := cli.EventGetAllMempool(context.Background(), &pb.Empty{})
		if err != nil {
			log.Errorf("setGRPCHandlers: cli.EventGetAllMempool: %s", err.Error())
			// return nil, err
		}

		for {
			mpRec, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Errorf("setGRPCHandlers: client.EventGetAllMempool: %s", err.Error())
			}

			mpRates := &mgo.Collection{}
			switch networtkID {
			case currencies.Main:
				mpRates = mempoolRates
			case currencies.Test:
				mpRates = mempoolRatesTest
			default:
				log.Errorf("setGRPCHandlers: wrong networkID:")
			}

			err = mpRates.Insert(store.MempoolRecord{
				Category: int(mpRec.Category),
				HashTX:   mpRec.HashTX,
			})
			if err != nil {
				log.Errorf("initGrpcClient: mpRates.Insert: %s", err.Error())
			}
		}
	}()

	// add transaction on every new tx on node
	go func() {
		stream, err := cli.EventAddMempoolRecord(context.Background(), &pb.Empty{})
		if err != nil {
			log.Errorf("setGRPCHandlers: cli.EventAddMempoolRecord: %s", err.Error())
			// return nil, err
		}

		mpRates := &mgo.Collection{}
		switch networtkID {
		case currencies.Main:
			mpRates = mempoolRates
		case currencies.Test:
			mpRates = mempoolRatesTest
		default:
			log.Errorf("setGRPCHandlers: wrong networkID:")
		}

		for {
			mpRec, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Errorf("setGRPCHandlers: client.EventAddMempoolRecord:stream.Recv: %s", err.Error())
			}
			err = mpRates.Insert(store.MempoolRecord{
				Category: int(mpRec.Category),
				HashTX:   mpRec.HashTX,
			})
			if err != nil {
				log.Errorf("initGrpcClient: mpRates.Insert: %s", err.Error())
			}
		}
	}()

	//deleting mempool record on block
	go func() {

		stream, err := cli.EventDeleteMempool(context.Background(), &pb.Empty{})
		if err != nil {
			log.Errorf("setGRPCHandlers: cli.EventGetAllMempool: %s", err.Error())
			// return nil, err
		}

		mpRates := &mgo.Collection{}
		switch networtkID {
		case currencies.Main:
			mpRates = mempoolRates
		case currencies.Test:
			mpRates = mempoolRatesTest
		default:
			log.Errorf("setGRPCHandlers: wrong networkID:")
		}

		for {
			mpRec, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Errorf("initGrpcClient: cli.EventDeleteMempool:stream.Recv: %s", err.Error())
			}

			query := bson.M{"hashtx": mpRec.Hash}
			err = mpRates.Remove(query)

			if err != nil {
				log.Errorf("setGRPCHandlers:mpRates.Remove: %s", err.Error())
			} else {
				log.Debugf("Tx removed: %s", mpRec.Hash)
			}
		}

	}()

	// new spendable output
	go func() {
		stream, err := cli.EventAddSpendableOut(context.Background(), &pb.Empty{})
		if err != nil {
			log.Errorf("setGRPCHandlers: cli.EventGetAllMempool: %s", err.Error())
		}

		spOutputs := &mgo.Collection{}
		spend := &mgo.Collection{}
		switch networtkID {
		case currencies.Main:
			spOutputs = spendableOutputs
			spend = spentOutputs
		case currencies.Test:
			spOutputs = spendableOutputsTest
			spend = spentOutputsTest
		default:
			log.Errorf("setGRPCHandlers: wrong networkID:")
		}

		for {
			gSpOut, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Errorf("initGrpcClient: cli.EventAddSpendableOut:stream.Recv: %s", err.Error())
			}

			query := bson.M{"userid": gSpOut.UserID, "txid": gSpOut.TxID, "address": gSpOut.Address}
			err = spend.Find(query).One(nil)

			if err == mgo.ErrNotFound {
				user := store.User{}
				sel := bson.M{"wallets.addresses.address": gSpOut.Address}
				err = usersData.Find(sel).One(&user)
				if err != nil && err != mgo.ErrNotFound {
					log.Errorf("SetWsHandlers: cli.On newIncomingTx: %s", err)
					return
				}
				spOut := generatedSpOutsToStore(gSpOut)

				log.Infof("Add spendable output : %v", gSpOut.String())

				exRates, err := GetLatestExchangeRate()
				if err != nil {
					log.Errorf("initGrpcClient: GetLatestExchangeRate: %s", err.Error())
				}
				spOut.StockExchangeRate = exRates

				query := bson.M{"userid": spOut.UserID, "txid": spOut.TxID, "address": spOut.Address}
				err = spOutputs.Find(query).One(nil)
				if err == mgo.ErrNotFound {
					//insertion
					err := spOutputs.Insert(spOut)
					if err != nil {
						log.Errorf("Create spOutputs:txsData.Insert: %s", err.Error())
					}
					continue
				}
				if err != nil && err != mgo.ErrNotFound {
					log.Errorf("Create spOutputs:spOutputs.Find %s", err.Error())
					continue
				}

				update := bson.M{
					"$set": bson.M{
						"txstatus": spOut.TxStatus,
					},
				}
				err = spOutputs.Update(query, update)
				if err != nil {
					log.Errorf("CreateSpendableOutputs:spendableOutputs.Update: %s", err.Error())
				}
			}

		}

	}()

	// delete spendable output
	go func() {
		stream, err := cli.EventDeleteSpendableOut(context.Background(), &pb.Empty{})
		if err != nil {
			log.Errorf("setGRPCHandlers: cli.EventGetAllMempool: %s", err.Error())
		}

		spOutputs := &mgo.Collection{}
		spend := &mgo.Collection{}
		switch networtkID {
		case currencies.Main:
			spOutputs = spendableOutputs
			spend = spentOutputs
		case currencies.Test:
			spOutputs = spendableOutputsTest
			spend = spentOutputsTest
		default:
			log.Errorf("setGRPCHandlers: wrong networkID:")
		}

		for {
			del, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Errorf("initGrpcClient: cli.EventDeleteMempool:stream.Recv: %s", err.Error())
			}

			i := 0
			for {
				//insert to spend collection
				err := spend.Insert(del)
				if err != nil {
					log.Errorf("DeleteSpendableOutputs:spend.Insert: %s", err)
				}

				query := bson.M{"userid": del.UserID, "txid": del.TxID, "address": del.Address}
				log.Infof("-------- query delete %v\n", query)
				err = spOutputs.Remove(query)
				if err != nil {
					log.Errorf("DeleteSpendableOutputs:spendableOutputs.Remove: %s", err.Error())
				} else {
					log.Infof("delete success √: %v", query)
					break
				}
				i++
				if i == 10 {
					break
				}
				time.Sleep(time.Second * 3)

			}

			log.Debugf("DeleteSpendableOutputs:spendableOutputs.Remove: %s", err)

		}
	}()

	// add to transaction history record and send ws notification on tx
	go func() {
		stream, err := cli.NewTx(context.Background(), &pb.Empty{})
		if err != nil {
			log.Errorf("setGRPCHandlers: cli.EventGetAllMempool: %s", err.Error())
		}

		for {
			gTx, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Errorf("initGrpcClient: cli.NewTx:stream.Recv: %s", err.Error())
			}
			tx := generatedTxDataToStore(gTx)

			setExchangeRates(&tx, gTx.Resync, tx.MempoolTime)
			setUserID(&tx)
			// setTxInfo(&tx)
			user := store.User{}
			// set wallet index and address index in input
			for i := 0; i < len(tx.WalletsInput); i++ {
				sel := bson.M{"wallets.addresses.address": tx.WalletsInput[i].Address.Address}
				err := usersData.Find(sel).One(&user)
				if err == mgo.ErrNotFound {
					continue
				} else if err != nil && err != mgo.ErrNotFound {
					log.Errorf("initGrpcClient: cli.On newIncomingTx: %s", err)
				}

				for _, wallet := range user.Wallets {
					for _, addr := range wallet.Adresses {
						if addr.Address == tx.WalletsInput[i].Address.Address {
							tx.WalletsInput[i].WalletIndex = wallet.WalletIndex
							tx.WalletsInput[i].Address.AddressIndex = addr.AddressIndex
						}
					}
				}
			}

			for i := 0; i < len(tx.WalletsOutput); i++ {
				sel := bson.M{"wallets.addresses.address": tx.WalletsOutput[i].Address.Address}
				err := usersData.Find(sel).One(&user)
				if err == mgo.ErrNotFound {
					continue
				} else if err != nil && err != mgo.ErrNotFound {
					log.Errorf("initGrpcClient: cli.On newIncomingTx: %s", err)
				}

				for _, wallet := range user.Wallets {
					for _, addr := range wallet.Adresses {
						if addr.Address == tx.WalletsOutput[i].Address.Address {
							tx.WalletsOutput[i].WalletIndex = wallet.WalletIndex
							tx.WalletsOutput[i].Address.AddressIndex = addr.AddressIndex
						}
					}
				}
			}

			log.Infof("New tx history in- %v out-%v\n", tx.WalletsInput, tx.WalletsOutput)

			err = saveMultyTransaction(tx, networtkID, gTx.Resync)
			if err != nil {
				log.Errorf("initGrpcClient: saveMultyTransaction: %s", err)
			}
			updateWalletAndAddressDate(tx, networtkID)
			sendNotifyToClients(tx, nsqProducer)

		}
	}()

	// watch for channel and push to node
	go func() {
		for {
			select {
			case addr := <-wa:
				a := addr
				rp, err := cli.EventAddNewAddress(context.Background(), &a)
				if err != nil {
					log.Errorf("NewAddressNode: cli.EventAddNewAddress %s\n", err.Error())
				}
				log.Debugf("EventAddNewAddress Reply %s", rp)

				rp, err = cli.EventResyncAddress(context.Background(), &pb.AddressToResync{
					Address: addr.Address,
				})
				if err != nil {
					log.Errorf("EventResyncAddress: cli.EventResyncAddress %s\n", err.Error())
				}
				log.Debugf("EventResyncAddress Reply %s", rp)

			}
		}
	}()

}
