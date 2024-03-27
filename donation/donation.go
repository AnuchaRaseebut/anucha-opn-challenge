package donation

import (
	"fmt"
	"log"
	"sort"
)

type Donation struct {
	Name           string
	AmountSubunits int64
	CCNumber       string
	CVV            string
	ExpMonth       int8
	ExpYear        int
}

type Summary struct {
	TotalReceived float64
	SuccessDonate float64
	FalseDonate   float64
	TotalDonors   int
	TopDonors     []Donor
}

type Donor struct {
	Name  string
	Total float64
}

func ProcessDonation(cfg Config, filepath string) (err error) {
	if donations, err := GetDonation(filepath); err != nil {
		log.Printf("Failed to get donation information: %v\n", err)
		return err
	} else {
		fmt.Println("performing donations...")
		return process(cfg, donations[1:])
	}
}

func process(cfg Config, donations []Donation) (err error) {
	summ := Summary{TotalDonors: len(donations)}
	for _, donate := range donations {
		summ.TotalReceived += float64(donate.AmountSubunits) / 100
		summ.TopDonors = append(summ.TopDonors, Donor{Name: donate.Name, Total: float64(donate.AmountSubunits) / 100})
		if err = createCharge(cfg, donate); err != nil {
			log.Printf("Failed to create charge for %s: %v\n", donate.Name, err)
			summ.FalseDonate += float64(donate.AmountSubunits) / 100
			continue
		} else {
			summ.SuccessDonate += float64(donate.AmountSubunits) / 100
		}
	}

	sort.SliceStable(summ.TopDonors, func(i, j int) bool {
		return summ.TopDonors[i].Total > summ.TopDonors[j].Total
	})

	summ.TopDonors = summ.TopDonors[:3]
	printSummary(summ)
	return nil
}

func printSummary(summ Summary) {
	avgPerPerson := summ.TotalReceived / float64(summ.TotalDonors)
	fmt.Println("done.")
	fmt.Printf("\n\ttotal received: THB %.2f\n", summ.TotalReceived)
	fmt.Printf("successfully donated: THB %.2f\n", summ.SuccessDonate)
	fmt.Printf("     faulty donation: THB %.2f\n", summ.FalseDonate)
	fmt.Printf("\n average per person: THB %.2f\n", avgPerPerson)
	fmt.Println("         top donors:")
	for _, donor := range summ.TopDonors {
		fmt.Println("                        ", donor.Name)
	}
}
