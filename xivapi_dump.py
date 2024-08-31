#!/usr/bin/python

"""
A command line tool to dump information about battle gear and weapons from the XIVAPI API to a local JSON file.
"""

import argparse
import csv
import json
from dataclasses import dataclass
from typing import Any, Dict

from frozendict import frozendict
import requests

gear_categories = [
    # "Shield",
    "Head",
    "Body",
    "Hands",
    "Legs",
    "Feet",
    "Necklace",
    "Earrings",
    "Bracelets",
    "Ring",
]

jobs = [
    # tanks
    # "GLA",
    "PLD",
    # "MRD",
    "WAR",
    # "DRK",
    "GNB",
    #  healers
    # "CNJ",
    "WHM",
    "SCH",
    "AST",
    "SGE",
    # melee dps
    # "LNC",
    "DRG",
    # "PGL",
    "MNK",
    # "ROG",
    "NIN",
    "SAM",
    "RPR",
    "VPR",
    # ranged physical dps
    # "ARC",
    "BRD",
    "MCH",
    "DNC",
    # ranged magical dps
    # "THM",
    "BLM",
    # "ACN",
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


limits = Limits(0, 0)


def init_argparse() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Download information about all FFXIV gear items from the official Eorzea Database"
    )

    parser.add_argument("-min", "--min_ilvl", default=690, type=int, nargs="?")
    parser.add_argument("-max", "--max_ilvl", default=999, type=int, nargs="?")
    parser.add_argument("-o", "--out", default="items-xivapi.json", nargs="?")
    parser.add_argument(
        "-f", "--format", choices=["json", "csv"], default="json", nargs="?"
    )

    return parser


def load(job: str, category: str) -> Dict[str, frozendict[str, Any]]:
    filters = [
        f"+LevelItem>={limits.min_ilvl}",
        f"+LevelItem<={limits.max_ilvl}",
        f"+ClassJobCategory.{job}=true",
        # f"ClassJobUse.ClassJobCategory.{job}=true",
        f'+ItemUICategory.Name="{category}"',
    ]
    query = " ".join(filters)
    limit = 3000
    fields = "Name,LevelEquip,LevelItem.value,DamagePhys,DamageMag,Delayms,BaseParam[].Name,BaseParamValue,CanBeHq,IsUnique,BaseParamSpecial[].Name,BaseParamValueSpecial,ItemUICategory.Name,EquipSlotCategory.value,MateriaSlotCount"

    query = requests.utils.quote(query)
    fields = requests.utils.quote(fields)

    url = f"https://beta.xivapi.com/api/1/search?sheets=Item&query={query}&limit={limit}&fields={fields}"

    # print(url)

    response = requests.get(url, timeout=30.0)

    items = map(lambda it: it["fields"], response.json()["results"])

    return {
        it["Name"]: {
            "Type": it["ItemUICategory"]["fields"]["Name"],
            "EquipSlotCategory": it["EquipSlotCategory"]["value"],
            "ItemLvl": it["LevelItem"]["value"],
            "EquipLvl": it["LevelEquip"],
            "PhysDmg": it["DamagePhys"],
            "MagDmg": it["DamageMag"],
            # "DelayMS": it["Delayms"],
            "MateriaSlotCount": it["MateriaSlotCount"],
            "CanBeHq": it["CanBeHq"],
            "IsUnique": it["IsUnique"],
            **dict(
                zip(
                    (p["fields"]["Name"] for p in it["BaseParam"]),
                    filter(lambda x: x > 0, it["BaseParamValue"]),
                )
            ),
            "BaseParamSpecial": dict(
                zip(
                    (p["fields"]["Name"] for p in it["BaseParamSpecial"]),
                    filter(lambda x: x > 0, it["BaseParamValueSpecial"]),
                )
            ),
        }
        for it in items
    }


def dump_gear_json(file_name: str):
    gear_index = {}
    gear = {}
    for job in jobs:
        print(f"dumping {job} gear")
        gear_index[job] = {}
        for cat in gear_categories:
            print(cat)
            items = load(job, cat)
            gear_index[job][cat] = [name for name in items]
            gear |= items

        print("Weapon")
        weapons = {}
        for weapon_type in weapon_categories[job]:
            weapons |= load(job, weapon_type)

        gear_index[job]["Weapon"] = [name for name in weapons]
        gear |= weapons

        print("OffHand")
        items = load(job, "Shield")
        gear_index[job]["OffHand"] = [name for name in items]
        gear |= items

    with open(file_name, "wt", encoding="utf-8") as file:
        json.dump({"index": gear_index, "items": gear}, file)


def dump_gear_csv(file_name: str):
    items = []
    for job in jobs:
        print(f"dumping {job} gear")
        for cat in gear_categories:
            print(cat)
            items.extend(load(job, cat).values())

        print("weapon")
        for weapon_type in weapon_categories[job]:
            items.extend(load(job, weapon_type).values())

    with open(file_name, "wt", encoding="utf-8", newline="") as file:
        keys = [
            "type",
            "name",
            "ilvl",
            "job level",
            "Physical Damage",
            "Magic Damage",
            "Delay",
            "materia slots",
            "Strength",
            "Dexterity",
            "Vitality",
            "Intelligence",
            "Mind",
            "Critical Hit",
            "Determination",
            "Direct Hit Rate",
            "Skill Speed",
            "Spell Speed",
            "Tenacity",
            "Piety",
        ]
        writer = csv.DictWriter(file, keys)
        writer.writeheader()
        writer.writerows(items)


def main():
    parser = init_argparse()
    args = parser.parse_args()
    limits.min_ilvl = args.min_ilvl
    limits.max_ilvl = args.max_ilvl

    print(f"loading gear from {limits.min_ilvl} to {limits.max_ilvl         }")

    if args.format == "json":
        dump_gear_json(args.out)
    elif args.format == "csv":
        dump_gear_csv(args.out)


if __name__ == "__main__":
    main()
