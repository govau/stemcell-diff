package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func getStemcellInfo(thing string) (map[string]string, error) {
	parts := strings.Split(thing, "-")
	if len(parts) != 2 {
		return nil, errors.New("should be format: xenial-97.18")
	}

	resp, err := http.Get(fmt.Sprintf("https://s3.amazonaws.com/bosh-aws-light-stemcells/%[1]s/light-bosh-stemcell-%[1]s-aws-xen-hvm-ubuntu-%[2]s-go_agent.tgz", parts[1], parts[0]))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	if err != nil {
		return nil, err
	}

	for {
		header, err := tr.Next()
		if err != nil {
			return nil, err
		}
		if header.Name != "packages.txt" {
			continue
		}
		if header.Size > (10 * 1024 * 1024) {
			return nil, errors.New("malformed header - packages file reported way too big")
		}
		packages := make([]byte, header.Size)
		_, err = io.ReadFull(tr, packages)
		if err != nil {
			return nil, err
		}

		rv := make(map[string]string)
		for _, line := range strings.Split(string(packages), "\n") {
			if !strings.HasPrefix(line, "ii") {
				continue
			}
			cols := strings.Fields(line)
			if len(cols) < 3 {
				return nil, errors.New("unrecognized row")
			}
			rv[cols[1]] = cols[2]
		}
		return rv, nil
	}
}

type tableWriter struct {
	rows [][]string
}

func (t *tableWriter) Write(cols []string) {
	t.rows = append(t.rows, cols)
}

func (t *tableWriter) Print(w io.Writer) {
	var colWidths []int
	for _, r := range t.rows {
		for ci, cv := range r {
			for ci >= len(colWidths) {
				colWidths = append(colWidths, 0)
			}
			if colWidths[ci] < len(cv) {
				colWidths[ci] = len(cv)
			}
		}
	}
	for _, r := range t.rows {
		for ci, cv := range r {
			fmt.Fprintf(w, "%-*s", colWidths[ci]+2, cv)

		}
		fmt.Fprint(w, "\n")
	}
}

func showMapDiff(nb, na string, b, a map[string]string, w io.Writer) {
	keys := make(map[string]interface{})
	for k := range a {
		keys[k] = nil
	}
	for k := range b {
		keys[k] = nil
	}
	keyList := make([]string, 0, len(keys))
	for k := range keys {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)

	tw := &tableWriter{}
	tw.Write([]string{"", nb, na})
	for _, k := range keyList {
		if a[k] != b[k] {
			tw.Write([]string{k, b[k], a[k]})
		}
	}
	tw.Print(w)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: stemcell-diff xenial-97.18 xenial-97.28")
	}

	before, err := getStemcellInfo(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	after, err := getStemcellInfo(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	showMapDiff(os.Args[1], os.Args[2], before, after, os.Stdout)
}
