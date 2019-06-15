package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sanandak/seg/seg2"
	"github.com/sanandak/seg/segy"
)

func main() {
	outfile := flag.String("o", "", "su output file name")
	nf := flag.Int("n", 1, "number of consecutive files to process")
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		log.Fatal("no input files specified")
	}
	// seg2 files are consecutive number nnn.dat
	// -n nf says process nnn, nnn+1...nnn+nf-1
	f0 := files[0]
	f0num, err := strconv.Atoi(strings.TrimSuffix(f0, filepath.Ext(f0)))
	//fmt.Println(files, *nf)
	if len(*outfile) == 0 {
		fnoex := strings.TrimSuffix(f0, filepath.Ext(f0))
		*outfile = fnoex + ".su"
	}

	fout, err := os.OpenFile(*outfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("err opening su outfile", *outfile)
	}
	defer fout.Close()
	for i := 0; i < *nf; i++ {
		fn := fmt.Sprintf("%d.dat", f0num+i)
		log.Println("processing file", i, fn)

		seg2trcs := seg2.ReadSEG2(fn)
		//fmt.Println(len(seg2trcs))
		//fmt.Println(seg2trcs[0].Data[:200])

		for _, t := range seg2trcs {
			segytrc := segy.Seg2Segy(t)
			//fmt.Println(len(segytrc))
			n, err := fout.Write(segytrc)
			if err != nil {
				log.Fatal("err writing su file", n, err)
			}
		}
	}
}
