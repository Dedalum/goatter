// cli.go
package cli

import (
    "flag"
    "fmt"
    "log"
    "os"
    "runtime"
    "strconv"

    "github.com/Dedalum/goatter/blockchain"
    "github.com/Dedalum/goatter/blockchain/wallet" // This is the new one
)

func (cli * CommandLine) Run(){
// big fat function here
}

func(cli *CommandLine) listAddresses() {
    wallets, _ := wallet.CreateWallets()
    addresses := wallets.GetAllAddresses()

    for _, address := range addresses {
        fmt.Println(address)
    }
}

func(cli *CommandLine) createWallet() {
    wallets, _ := wallet.CreateWallets()
    address := wallets.AddWallet()
    wallets.SaveFile()

    fmt.Printf("New address is: %s\n", address)
}

func (cli *CommandLine) Run() {
    cli.validateArgs()

    getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
    createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
    sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
    createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError) // this Cmd is new
    listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError) // this Cmd is new

    getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
    createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
    sendFrom := sendCmd.String("from", "", "Source wallet address")
    sendTo := sendCmd.String("to", "", "Destination wallet address")
    sendAmount := sendCmd.Int("amount", 0, "Amount to send")

    switch os.Args[1] {
    case "getbalance":
        err := getBalanceCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "createblockchain":
        err := createBlockchainCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "printchain":
        err := printChainCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "send":
        err := sendCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "listaddresses": // this case statement is new
        err := listAddressesCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "createwallet": // this case statement is new
        err := createWalletCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    default:
        cli.printUsage()
        runtime.Goexit()
    }

    if getBalanceCmd.Parsed() {
        if *getBalanceAddress == "" {
            getBalanceCmd.Usage()
            runtime.Goexit()
        }
        cli.getBalance(*getBalanceAddress)
    }

    if createBlockchainCmd.Parsed() {
        if *createBlockchainAddress == "" {
            createBlockchainCmd.Usage()
            runtime.Goexit()
        }
        cli.createBlockChain(*createBlockchainAddress)
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }

    if sendCmd.Parsed() {
        if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
            sendCmd.Usage()
            runtime.Goexit()
        }

        cli.send(*sendFrom, *sendTo, *sendAmount)
    }
    if listAddressesCmd.Parsed() { // this if statement is new
        cli.listAddresses()
    }
    if createWalletCmd.Parsed(){ // this if statement is new
        cli.createWallet()
    }
}