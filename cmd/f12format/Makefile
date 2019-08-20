test:
	cd ../gosk && go build
	cd ../../ && ./cmd/gosk/gosk testdata/03_day_harib00i_ipl10.nas ipl10.bin
	go build
	./f12format -f 1440 -B ../../ipl10.bin -i ../../output.img
