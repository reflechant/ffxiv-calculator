#!/usr/bin/python

"""
A Scrapy-based web crawler for the official FFXIV Eorzea Database. Generates a JSON file with gear information.
"""

import re
import argparse

import scrapy
from scrapy.crawler import CrawlerProcess


# [category2, category3] are numbers for the URL
categories = {
    "Arms": [1, ""],  # for all classes
    "Shield": [3, 11],
    "Head": [3, 34],
    "Body": [3, 35],
    "Hands": [3, 37],
    "Legs": [3, 36],
    "Feet": [3, 38],
    "Earring": [4, 41],
    "Necklace": [4, 41],
    "Bracelets": [4, 42],
    "Ring": [4, 43],
}

limits = {
    "min_ilvl": 1,
    "max_ilvl": 999,
}


class ItemSpider(scrapy.Spider):
    name = "item_spider"

    start_urls = ["https://eu.finalfantasyxiv.com/lodestone/playguide/db/"]

    def parse_item(self, response, cat_name):
        name: str = response.css("h2.db-view__item__text__name::text").get()
        name = name.replace("\n", "")
        name = name.replace("\t", "")

        ilvl = int(response.css("div.db-view__item_level::text").get().split()[-1])

        spec_names = response.css(
            "div.db-view__item_spec > div > div.db-view__item_spec__name::text"
        ).getall()
        spec_values = response.css(
            "div.db-view__item_spec > div > div.db-view__item_spec__value > strong::text"
        ).getall()
        specs = zip(spec_names, (float(v) for v in spec_values))

        job = response.css("div.db-view__item_equipment__class::text").get().split()
        job_lvl_str = response.css("div.db-view__item_equipment__level::text").get()
        job_lvl_regex = re.compile(r"Lv\.\s(\d+)")
        job_lvl = int(job_lvl_regex.match(job_lvl_str).group(1))

        stat_bonus_regexp = re.compile(r".*\+(\d+)")
        bonus_els = response.css("ul.db-view__basic_bonus > li")
        bonuses = {
            el.css("span::text").get(): int(stat_bonus_regexp.match(el.get()).group(1))
            for el in bonus_els
        }

        materia_slots = len(
            response.css("ul.db-view__materia_socket > li.socket, normal")
        )

        yield {
            "type": cat_name,
            "name": name,
            "ilvl": ilvl,
            "job": job,
            "job level": job_lvl,
            **dict(specs),
            **bonuses,
            "materia slots": materia_slots,
        }

    def parse_category(self, response, cat_name):
        for item in response.css("a.db-table__txt--detail_link"):
            item_url = item.css("::attr(href)").get()
            item_url = response.urljoin(item_url)
            yield scrapy.Request(
                url=item_url, callback=self.parse_item, cb_kwargs={"cat_name": cat_name}
            )

        next_page = response.css('li.next a::attr("href")').get()
        if next_page is not None:
            next_page = response.urljoin(next_page)
            yield scrapy.Request(
                url=next_page,
                callback=self.parse_category,
                cb_kwargs={"cat_name": cat_name},
            )

    def parse(self, response):
        for name, cats in categories.items():
            cat2, cat3 = cats
            cat_url = f"https://eu.finalfantasyxiv.com/lodestone/playguide/db/item/?patch=&db_search_category=item&category2={cat2}&category3={cat3}&difficulty=&min_item_lv={limits['min_ilvl']}&max_item_lv={limits['max_ilvl']}&min_gear_lv=&max_gear_lv=&min_craft_lv=&max_craft_lv=&q="
            yield scrapy.Request(
                url=cat_url, callback=self.parse_category, cb_kwargs={"cat_name": name}
            )


def init_argparse() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        description="Download information about all FFXIV gear items from the official Eorzea Database"
    )

    parser.add_argument("-min", "--min_ilvl", default=1, type=int)
    parser.add_argument("-max", "--max_ilvl", default=999, type=int)
    parser.add_argument("-o", "--out", default="items.json")

    return parser


def main():
    parser = init_argparse()
    args = parser.parse_args()
    limits["min_ilvl"] = args.min_ilvl
    limits["max_ilvl"] = args.max_ilvl

    settings = {
        "FEEDS": {
            args.out: {"format": "json"},
        },
        "LOG_LEVEL": "WARNING",
    }
    runner = CrawlerProcess(settings)

    runner.crawl(ItemSpider)
    runner.start()


if __name__ == "__main__":
    main()
