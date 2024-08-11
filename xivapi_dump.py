#!/usr/bin/python

"""
A command line tool to dump information about battle gear and weapons from the XIVAPI API to a local JSON file.
"""

import argparse
import json
from dataclasses import dataclass
from typing import List, Dict, Any

import requests

gear_categories = [
    "Shield",
    "Head",
    "Body",
    "Legs",
    "Hands",
    "Feet",
    "Necklace",
    "Earrings",
    "Bracelets",
    "Ring",
]

jobs = [
    # tanks
    "GLA",
    "PLD",
    "MRD",
    "WAR",
    "DRK",
    "GNB",
    #  healers
    "CNJ",
    "WHM",
    "SCH",
    "AST",
    "SGE",
    # melee dps
    "LNC",
    "DRG",
    "PGL",
    "MNK",
    "ROG",
    "NIN",
    "SAM",
    "RPR",
    "VPR",
    # ranged physical dps
    "ARC",
    "BRD",
    "MCH",
    "DNC",
    # ranged magical dps
    "THM",
    "BLM",
    "ACN",
    "SMN",
    "RDM",
    "PCT",
    # limited
    "BLU",
]

weapon_categories = {
    #  tanks
    "GLA": ["Gladiator's Arm"],
    "PLD": ["Gladiator's Arm"],
    "MRD": ["Marauder's Arm"],
    "WAR": ["Marauder's Arm"],
    "DRK": ["Dark Knight's Arm"],
    "GNB": ["Gunbreaker's Arm"],
    # healers
    "CNJ": ["One–handed Conjurer's Arm", "Two–handed Conjurer's Arm"],
    "WHM": ["One–handed Conjurer's Arm", "Two–handed Conjurer's Arm"],
    "SCH": ["Scholar's Arm"],
    "AST": ["Astrologian's Arm"],
    "SGE": ["Sage's Arm"],
    # melee DPS
    "LNC": ["Lancer's Arm"],
    "DRG": ["Lancer's Arm"],
    "PGL": ["Pugilist's Arm"],
    "MNK": ["Pugilist's Arm"],
    "ROG": ["Rogue's Arm"],
    "NIN": ["Rogue's Arm"],
    "SAM": ["Samurai's Arm"],
    "RPR": ["Reaper's Arm"],
    "VPR": ["Viper's Arm"],
    # ranged physical DPS
    "ARC": ["Archer's Arm"],
    "BRD": ["Archer's Arm"],
    "MCH": ["Machinist's Arm"],
    "DNC": ["Dancer's Arm"],
    # ranged magical DPS
    "THM": ["One–handed Thaumaturge's Arm", "Two–handed Thaumaturge's Arm"],
    "BLM": ["One–handed Thaumaturge's Arm", "Two–handed Thaumaturge's Arm"],
    "ACN": ["Arcanist's Grimoire"],
    "SMN": ["Arcanist's Grimoire"],
    "RDM": ["Red Mage's Arm"],
    "PCT": ["Pictomancer's Arm"],
    # limited
    "BLU": ["Blue Mage's Arm"],
}


@dataclass
class Limits:
    min_ilvl: int
    max_ilvl: int


limits = Limits(690, 999)


def init_argparse() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Download information about all FFXIV gear items from the official Eorzea Database"
    )

    parser.add_argument("-min", "--min_ilvl", default=710, type=int)
    parser.add_argument("-max", "--max_ilvl", default=999, type=int)
    parser.add_argument("-o", "--out", default="items-xivapi.json")

    return parser


def load(job: str, category: str) -> Dict[str, Dict[str, Any]]:
    filters = [
        f"LevelItem>={limits.min_ilvl}",
        f"LevelItem<={limits.max_ilvl}",
        f"ClassJobCategory.{job}=true",
        f"ItemUICategory.Name=\"{category}\"",
    ]
    query = " ".join(filters)
    limit = 3000
    fields = "Name,LevelEquip,LevelItem.value,DamagePhys,DamageMag,Delayms,BaseParam[].Name,BaseParamValue,ItemUICategory.Name,MateriaSlotCount"

    query = requests.utils.quote(query)
    fields = requests.utils.quote(fields)

    url = f"https://beta.xivapi.com/api/1/search?sheets=Item&query={query}&limit={limit}&fields={fields}"

    response = requests.get(url, timeout=30.0)

    items = map(lambda it: it['fields'],response.json()["results"])

    return {
        it["Name"]: {
            "ilvl": it["LevelItem"]["value"],
            "job level": it["LevelEquip"],
            "Physical Damage": it["DamagePhys"],
            "Delay": it["Delayms"],
            "materia slots": it["MateriaSlotCount"],
        }
        | dict(
            zip(
                (p["fields"]["Name"] for p in it["BaseParam"]),
                filter(lambda x: x > 0, it["BaseParamValue"]),
            )
        )
        for it in items
    }


def dump_gear(file_name: str):
    gear = {}
    for job in jobs:
        print(f"dumping {job} gear")
        gear[job] = {}
        for cat in gear_categories:
            print(cat)
            gear[job][cat] = load(job, cat)

        gear[job]["weapon"] = {}
        print("weapon")
        for weapon_type in weapon_categories[job]:
            gear[job]["weapon"] |= load(job, weapon_type)

    with open(file_name, "wt", encoding="utf-8") as file:
        json.dump(gear, file)


def main():
    parser = init_argparse()
    args = parser.parse_args()
    limits.min_ilvl = args.min_ilvl
    limits.max_ilvl = args.max_ilvl
    dump_gear(args.out)


if __name__ == "__main__":
    main()
