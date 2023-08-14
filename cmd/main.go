package main

import 

func main(){

	//И
	config.ConfigSetup()
	dbObject := db.NewDB()
	csh := db.NewCache(dbObject)

	// Запуск сервера для выдачи OrderOut

	myApi := api.NewApi(csh)


	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func(){
		for range signalChan{
			fmt.Printf("\nRecived an interrupt, unsubscribing and...")
			csh.Finish()
			sh.Finish()
			myApi.Finish()

			cleanupDone <- true
		}
	}()
	<-cleanupDone

}