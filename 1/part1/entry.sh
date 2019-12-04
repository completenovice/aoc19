#!/bin/bash

fuel=0

fuelCalculator () {
	interim=$(($1 / 3))
	floored=$(echo $interim | perl -nl -MPOSIX -e 'print floor($_);')
	required=$(($floored - 2))

	if [[ $required -ge 0 ]]
	then
		return $required
	fi
	return 0
}

while read mass
do
	ARRAY=()
	echo "MASS: $mass"
	while [[ $mass -gt 0 ]]
	do
		mass=fuelCalculator "$mass"
		echo "REQUIRES: $mass"
		ARRAY+=$required
	done

	fuel=$(($fuel + $required))
done < input.txt


echo $fuel
