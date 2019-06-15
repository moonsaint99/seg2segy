# seg2segy

command line program to convert SEG-2 files to SU (seismic unix) pseudo SEG-Y format. Written in go (htts://golang.org)

Usage: seg2segy [-o outfile] [-n nfiles] nnn.dat [mmm.dat ...]

Either specify a list of one or more seg2 files (nnn.dat mmm.dat ...) or specify one file (`nnn.dat`) and `-n nfiles` to process consecutive files (nnn, nnn+1 ... nnn+nfiles-1)

By default outfile is `nnn.su` unless `-o outfile` is specified.

Many data acquisition systems (e.g., Geometrics Geode) generate consecutively numbered files.

Seismic Unix: https://cwp.mines.edu/software  
https://github.com/JohnWStockwellJr/SeisUnix  

# Install

  - go get github.com/sanandak/seg2segy
  - cd <seg2segy> dir
  - go get github.com/sanandak/seg
  - go build
  - go install