package methods

import (
	"os"
	"strconv"
)

var validPincode = make(map[string]string)

var pincodes []string
var locality []string

func init() {

	ch := make(chan bool)
	var a, b bool

	a, b = false, false

	runPin, _ := strconv.Atoi(os.Args[2])

	if runPin == 0 {
		a = true
	} else {
		go func() {
			initA()
			ch <- true
		}()
	}

	locality = []string{
		"A F station yelahanka",
		"Adugodi",
		"Agara",
		"Agram",
		"Air Force hospital",
		"Amruthahalli",
		"Anandnagar",
		"Anekal",
		"Anekalbazar",
		"Arabic College",
		"Aranya Bhavan",
		"Ashoknagar",
		"Attibele",
		"Attur",
		"Austin Town",
		"Avalahalli",
		"Avani Sringeri mutt",
		"Avenue Road",
		"B Sk ii stage",
		"Bsf Campus yelahanka",
		"Bagalgunte",
		"Bagalur",
		"Balepete",
		"Banashankari",
		"Banashankari Iii stage",
		"Banaswadi",
		"Bandikodigehalli",
		"Bangalore Air port",
		"Bangalore Bazaar",
		"Bangalore City",
		"Bangalore Corporation building",
		"Bangalore Dist offices bldg",
		"Bangalore Fort",
		"Bangalore Sub fgn post",
		"Bangalore",
		"Bannerghatta",
		"Bannerghatta Road",
		"Bapujinagar",
		"Basavanagudi",
		"Basavaraja Market",
		"Basaveshwaranagar",
		"Basaveswaranagar II stage",
		"Bellandur",
		"Benson Town",
		"Bestamaranahalli",
		"Bettahalsur",
		"Bhashyam Circle",
		"Bhattarahalli",
		"Bidaraguppe",
		"Bidrahalli",
		"Bnagalore Viswavidalaya",
		"Bommanahalli",
		"Brigade Road",
		"Byatarayanapura",
		"C.V.raman nagar",
		"Cmp Centre and school",
		"Crpf Campus yelahanka",
		"Cahmrajendrapet",
		"Chamrajpet",
		"Chamrajpet Bazar",
		"Chandra Lay out",
		"Chickpet",
		"Chikkabettahalli",
		"Chikkajala",
		"Chikkalasandra",
		"Chikkanahalli",
		"Chunchanakuppe",
		"Cubban Road",
		"Dasarahalli",
		"Deepanjalinagar",
		"Devanagundi",
		"Devarjeevanahalli",
		"Devasandra",
		"Dharmaram College",
		"Doddagubbi",
		"Doddajala",
		"Doddakallasandra",
		"Doddanekkundi",
		"Domlur",
		"Dommasandra",
		"Doorvaninagar",
		"Dr. ambedkar veedhi",
		"Electronics City",
		"Fraser Town",
		"G.K.v.k.",
		"Gaviopuram Extension",
		"Gaviopuram Guttanahalli",
		"Gayathrinagar",
		"Girinagar",
		"Goraguntepalya",
		"Goripalya",
		"Governmemnt Electric factory",
		"Govindapalya",
		"Gunjur",
		"H M t",
		"H.A. farm",
		"H.A.l ii stage",
		"H.K.p. road",
		"Hsr Layout	",
		"Hampinagar",
		"Handenahalli",
		"Harogadde",
		"Hebbal Kempapura",
		"Hennagara",
		"Highcourt",
		"Hongasandra",
		"Hoodi",
		"Horamavu",
		"Hosakerehalli",
		"Hosur Road",
		"Hulimangala",
		"Hulimavu",
		"Hulsur Bazaar",
		"Hunasamaranahalli",
		"Isro Anthariksha bhavan",
		"Immedihalli",
		"Indalavadi	",
		"Indiranagar",
		"Indiranagar Com. complex",
		"Industrial Estate",
		"Ittamadu Layout",
		"J P nagar",
		"J.C.nagar",
		"Jakkur",
		"Jalahalli",
		"Jalahalli East",
		"Jalahalli Village",
		"Jalahalli West",
		"Jalavayuvihar",
		"Jayanagar",
		"Jayanagar West",
		"Jayangar East",
		"Jayangar Iii block",
		"Jeevanahalli",
		"Jeevanbhimanagar",
		"Jigani",
		"Jp Nagar iii phase",
		"K H b colony",
		"K. g. road",
		"K.P.west",
		"Kacharakanahalli",
		"Kadabagere",
		"Kadugodi",
		"Kalkunte",
		"Kalyanagar",
		"Kamagondanahalli",
		"Kamakshipalya",
		"Kannamangala",
		"Kannur",
		"Kanteeravanagar",
		"Kathriguppe",
		"Kenchanahalli",
		"Kendriya Sadan",
		"Kendriya Vihar",
		"Kodigehalli",
		"Konanakunte",
		"Koramangala",
		"Koramangala I block",
		"Koramangala Vi bk",
		"Kothanur",
		"Krishnarajapuram",
		"Krishnarajapuram R s",
		"Kugur",
		"Kumaraswamy Layout",
		"Kumbalgodu",
		"Kundalahalli",
		"Lalbagh West",
		"Legislators Home",
		"Lingarajapuram",
		"M S r road",
		"Madhavan Park",
		"Madivala",
		"Magadi Road",
		"Mahadevapura",
		"Mahalakshmipuram Layout",
		"Mahatma Gandhi road",
		"Malkand Lines",
		"Mallathahalli",
		"Malleswaram",
		"Malleswaram West",
		"Mandalay Lines",
		"Marathahalli Colony",
		"Marsur",
		"Maruthi Sevanagar",
		"Mathikere",
		"Mavalli",
		"Mayasandra",
		"Medihalli",
		"Medimallasandra",
		"Mico Layout",
		"Milk Colony",
		"Mount St joseph",
		"Msrit",
		"Mundur",
		"Museum Road",
		"Muthanallur",
		"Muthusandra",
		"Nal",
		"Naduvathi",
		"Nagarbhavi",
		"Nagasandra",
		"Nagavara",
		"Nandinilayout",
		"Narasimharaja Colony",
		"Narasimjharaja Road",
		"Narayan Pillai street",
		"Nayandahalli",
		"Neralur",
		"New Tharaggupet",
		"New Thippasandra",
		"Okalipuram	",
		"P&t Col. kavalbyrasandra",
		"Padmanabhnagar",
		"Palace Guttahalli",
		"Panathur",
		"Pasmpamahakavi Road",
		"Peenya I stage	",
		"Peenya Ii stage",
		"Peenya Small industries",
		"R T nagar",
		"R.M.v. extension ii stage",
		"Rajajinagar",
		"Rajajinagar I block",
		"Rajajinagar Ivth block",
		"Rajanakunte",
		"Rajarajeshwarinagar",
		"Rajbhavan",
		"Ramachandrapuram",
		"Ramagondanahalli",
		"Ramakrishna Hegde nagar",
		"Ramamurthy Nagar",
		"Rameshnagar",
		"Richmond Town",
		"Rv Niketan",
		"Sadashivanagar",
		"Sahakaranagar P.o",
		"Samandur",
		"Samethanahalli",
		"Sampangiramnagar",
		"Sarjapura",
		"Science Institute",
		"Seshadripuram",
		"Shankarpura",
		"Shanthinagar",
		"Sidihoskote",
		"Singanayakanahalli",
		"Sivan Chetty gardens",
		"Someswarapura",
		"Sri Jayachamarajendra road",
		"Srirampuram",
		"St. john's medical college",
		"St. thomas town",
		"State Bank of mysore colony",
		"Subhashnagar",
		"Subramanyapura",
		"Swimming Pool extn",
		"Tarahunise",
		"Tavarekere	",
		"Taverekere	",
		"Thambuchetty Palya",
		"Thammanayakanahalli",
		"Tilaknagar",
		"Training Command iaf",
		"Tyagrajnagar",
		"Ullalu Upanagara",
		"Vanakanahalli",
		"Vartur",
		"Vasanthnagar",
		"Venkatarangapura",
		"Venkateshapura",
		"Vidhana Soudha",
		"Vidyanagara",
		"Vidyaranyapura",
		"Vijayanagar",
		"Vijayanagar East",
		"Vikramnagar",
		"Vimapura",
		"Virgonagar",
		"Visveswarapuram",
		"Viswaneedam",
		"Vittalnagar",
		"Viveknagar",
		"Vyalikaval Extn",
		"Wheel And axle plant",
		"Whitefield",
		"Wilson Garden",
		"Yadavanahalli",
		"Yediyur",
		"Yelachenahalli",
		"Yelahanka",
		"Yelahanka Satellite town",
		"Yemalur",
		"Yeshwanthpur Bazar",
		"Yeswanthpura",
	}

	pincodes = []string{
		"560063",
		"560030",
		"560034",
		"560007",
		"560007",
		"560092",
		"560024",
		"562106",
		"562106",
		"560045",
		"560003",
		"560050",
		"562107",
		"560064",
		"560047",
		"560026",
		"560086",
		"560002",
		"560070",
		"560064",
		"560073",
		"562149",
		"560053",
		"560050",
		"560085",
		"560043",
		"562149",
		"560017",
		"560001",
		"560002",
		"560002",
		"560009",
		"560002",
		"560025",
		"560001",
		"560083",
		"560076",
		"560026",
		"560004",
		"560002",
		"560079",
		"560086",
		"560103",
		"560046",
		"562106",
		"562157",
		"560010",
		"560049",
		"562107",
		"560049",
		"560056",
		"560068",
		"560001",
		"560026",
		"560093",
		"560025",
		"560064",
		"560002",
		"560018",
		"560018",
		"560040",
		"560053",
		"560097",
		"562157",
		"560061",
		"562130",
		"562130",
		"560001",
		"560050",
		"560026",
		"560067",
		"560045",
		"560036",
		"560029",
		"562149",
		"562157",
		"560062",
		"560037",
		"560071",
		"562125",
		"560016",
		"560001",
		"560100",
		"560005",
		"560065",
		"560019",
		"560019",
		"560021",
		"560085",
		"560022",
		"560026",
		"560026",
		"560013",
		"560087",
		"560013",
		"560024",
		"560008",
		"560051",
		"560102",
		"560104",
		"562125",
		"562106",
		"560024",
		"562106",
		"560001",
		"560068",
		"560048",
		"560043",
		"560085",
		"560030",
		"562106",
		"560076",
		"560008",
		"562157",
		"560094",
		"560066",
		"562106",
		"560038",
		"560038",
		"560010",
		"560085",
		"560078",
		"560006",
		"560064",
		"560013",
		"560014",
		"560013",
		"560015",
		"560043",
		"560041",
		"560070",
		"560069",
		"560011",
		"560005",
		"560066",
		"562106",
		"560078",
		"560079",
		"560009",
		"560020",
		"560084",
		"562130",
		"560067",
		"560067",
		"560043",
		"560015",
		"560079",
		"560067",
		"562149",
		"560096",
		"560085",
		"560098",
		"560034",
		"560064",
		"560092",
		"560062",
		"560034",
		"560034",
		"560095",
		"560077",
		"560036",
		"560016",
		"562125",
		"560078",
		"560074",
		"560037",
		"560004",
		"560001",
		"560084",
		"560054",
		"560011",
		"560068",
		"560023",
		"560048",
		"560086",
		"560001",
		"560033",
		"560056",
		"560003",
		"560055",
		"560033",
		"560037",
		"562106",
		"560033",
		"560054",
		"560004",
		"562107",
		"560049",
		"560067",
		"560076",
		"560055",
		"560076",
		"560054",
		"560049",
		"560025",
		"560099",
		"560087",
		"560017",
		"560067",
		"560072",
		"560073",
		"560045",
		"560096",
		"560019",
		"560002",
		"560001",
		"560039",
		"562107",
		"560002",
		"560075",
		"560021",
		"560032",
		"560070",
		"560003",
		"560087",
		"560004",
		"560058",
		"560058",
		"560058",
		"560032",
		"560094",
		"560010",
		"560010",
		"560010",
		"560064",
		"560098",
		"560001",
		"560021",
		"560066",
		"560045",
		"560016",
		"560037",
		"560025",
		"560059",
		"560080",
		"560092",
		"562106",
		"560067",
		"560027",
		"562125",
		"560012",
		"560020",
		"560004",
		"560027",
		"562106",
		"560064",
		"560042",
		"560008",
		"560002",
		"560021",
		"560034",
		"560084",
		"560050",
		"560009",
		"560061",
		"560003",
		"562157",
		"562130",
		"560029",
		"560049",
		"562106",
		"560041",
		"560006",
		"560028",
		"560056",
		"562106",
		"560087",
		"560052",
		"560003",
		"560045",
		"560001",
		"562157",
		"560097",
		"560040",
		"560040",
		"560078",
		"560017",
		"560049",
		"560004",
		"560091",
		"560018",
		"560047",
		"560003",
		"560064",
		"560066",
		"560027",
		"562107",
		"560070",
		"560078",
		"560064",
		"560064",
		"560037",
		"560022",
		"560022",
	}

	gos := 1

	proc := make(chan bool, gos)
	allDone := make(chan bool)

	for i := 0; i < gos; i++ {
		go func(n int) {
			for j := n; j < len(locality); j += gos {
				_, x := validPincode[locality[j]]

				if !x {
					validPincode[locality[j]] = pincodes[j]
				}
			}
			proc <- true
		}(i)
	}

	go func() {
		for i := 0; i < gos; i++ {
			<-proc
		}
		allDone <- true
	}()

	func() {
		for {
			select {
			case <-allDone:
				b = true
				if a && b {
					return
				}
			case <-ch:
				a = true
				if a && b {
					return
				}
			}
		}
	}()
}

func IsPresent(pincode string, locality string) (check bool) {
	var temp string
	check = true
	temp, t := validPincode[locality]

	if t && (temp == pincode) {
		check = false
	}

	return
}
