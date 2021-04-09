package quic

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"sync"

	quic "github.com/lucas-clemente/quic-go"
)

func client(addr string) error {
	ctx, cancel := context.WithCancel(context.Background())

	config := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quicssh"},
	}

	log.Printf("Dialing %q...", addr)
	session, err := quic.DialAddr(addr, config, nil)
	if err != nil {
		return err
	}

	log.Printf("Opening stream sync...")
	stream, err := session.OpenStreamSync(ctx)
	if err != nil {
		return err
	}

	log.Printf("Piping stream with QUIC...")
	var wg sync.WaitGroup
	wg.Add(3)
	c1 := readAndWrite(ctx, stream, os.Stdout, &wg)
	c2 := readAndWrite(ctx, os.Stdin, stream, &wg)
	select {
	case err = <-c1:
		if err != nil {
			return err
		}
	case err = <-c2:
		if err != nil {
			return err
		}
	}
	cancel()
	wg.Wait()
	return nil
}
