#!/bin/sh

for i in trainstation2.com portcity-shiptycoon.com puzzleadventuregame.com candyincgame.com; do
    mkdir -p tests/${i}/client-resources/
done

# DA2
touch tests/puzzleadventuregame.com/client-resources/client-data-000.558.sqlite
touch tests/puzzleadventuregame.com/client-resources/client-data-000.560.sqlite
touch tests/puzzleadventuregame.com/client-resources/client-data-latest.sqlite
mkdir tests/puzzleadventuregame.com/locations-000.558
mkdir tests/puzzleadventuregame.com/locations-000.560
mkdir tests/puzzleadventuregame.com/locations-latest

# TS2
touch tests/trainstation2.com/client-resources/client-data-204.009.sqlite
touch tests/trainstation2.com/client-resources/client-data-204.010.sqlite
touch tests/trainstation2.com/client-resources/client-data-204.011.sqlite
touch tests/trainstation2.com/client-resources/client-data-204.012.sqlite
touch tests/trainstation2.com/client-resources/client-data-204.013.sqlite
touch tests/trainstation2.com/client-resources/client-data-city_loop_2020_11-204.009.sqlite
touch tests/trainstation2.com/client-resources/client-data-featured-latest.sqlite
touch tests/trainstation2.com/client-resources/client-data-latest.sqlite
touch tests/trainstation2.com/client-resources/client-data-offer_badges_2022-204.013.sqlite
touch tests/trainstation2.com/client-resources/client-data-region_progression_2022_08-204.009.sqlite
touch tests/trainstation2.com/client-resources/client-data-region_progression_2022_08-204.010.sqlite
touch tests/trainstation2.com/client-resources/client-data-region_progression_2022_08-204.011.sqlite
touch tests/trainstation2.com/client-resources/client-data-region_progression_2022_08-204.012.sqlite
touch tests/trainstation2.com/client-resources/client-data-region_progression_2022_08-204.013.sqlite