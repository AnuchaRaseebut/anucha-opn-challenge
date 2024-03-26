package donation

import (
	"anucha-challenge/challenge-go/cipher"
	"bufio"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetDonation(filepath string) (donations []Donation, err error) {
	decryptPath := strings.TrimSuffix(filepath, ".rot128")
	encryptedFile, err := os.Open(filepath)
	if err != nil {
		log.Printf("Failed to open encrypt file: %v\n", err)
		return
	}
	defer encryptedFile.Close()

	rotReader, err := cipher.NewRot128Reader(encryptedFile)
	if err != nil {
		log.Printf("Failed to create Rot128Reader: %v\n", err)
		return
	}

	decryptedFile, err := os.Create(decryptPath)
	if err != nil {
		fmt.Printf("Failed to create decrypted file: %v\n", err)
		os.Exit(1)
	}
	defer decryptedFile.Close()

	writer := bufio.NewWriter(decryptedFile)
	defer writer.Flush()

	_, err = io.Copy(writer, rotReader)
	if err != nil {
		fmt.Printf("Failed to decrypt file: %v\n", err)
		os.Exit(1)
	}

	// Read the decrypted file
	decryptedFile.Seek(0, 0)
	csvReader := csv.NewReader(decryptedFile)
	rows, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read decrypted file: %v\n", err)
		os.Exit(1)
	}

	for _, row := range rows {
		amountSubunits, _ := strconv.Atoi(row[1])
		expMonthInt, _ := strconv.Atoi(row[4])
		expYearInt, _ := strconv.Atoi(row[5])

		donation := Donation{
			Name:           row[0],
			AmountSubunits: int64(amountSubunits),
			CCNumber:       row[2],
			CVV:            row[3],
			ExpMonth:       int8(expMonthInt),
			ExpYear:        expYearInt,
		}

		donations = append(donations, donation)
	}
	return
}
