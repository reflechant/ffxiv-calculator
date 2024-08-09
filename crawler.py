#!/usr/bin/python

"""
A Scrapy-based web crawler for the official FFXIV Eorzea Database to generate JSONL files with gear to use in gearset calculation
"""

import re

import scrapy
from scrapy.crawler import CrawlerProcess
from scrapy.utils.project import get_project_settings
from scrapy.utils.log import configure_logging


# [category2, category3] are numbers for the URL
categories = {
    "Arms": [1, ""],
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


class WeaponSpider(scrapy.Spider):
    name = "weapon_spider"
    min_ilvl = 710

    start_urls = [
        f"https://eu.finalfantasyxiv.com/lodestone/playguide/db/item/?patch=&db_search_category=item&category2=1&category3=106&difficulty=&min_item_lv={min_ilvl}&max_item_lv=&min_gear_lv=&max_gear_lv=&min_craft_lv=&max_craft_lv=&q="
    ]

    def parse_weapon(self, response):
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

        dmg = int(response.css("div.sys_nq_element > div > strong::text").get())

        job = response.css("div.db-view__item_equipment__class::text").get()
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
            "name": name,
            "ilvl": ilvl,
            "dmg": dmg,
            "job": job,
            "job level": job_lvl,
            **dict(specs),
            **bonuses,
            "materia slots": materia_slots,
        }

    def parse(self, response):
        for weapon in response.css("a.db-table__txt--detail_link"):
            a = weapon.css("::attr(href)").get()
            yield response.follow(a, callback=self.parse_weapon)

        next_page = response.css('li.next a::attr("href")').get()
        if next_page is not None:
            yield response.follow(next_page, self.parse)


def main():
    # settings = get_project_settings()
    # settings.set("LOG_LEVEL", "WARNING", priority="spider")
    settings = {
        "FEEDS": {
            "items.json": {"format": "json"},
        },
    }
    # configure_logging(settings)
    runner = CrawlerProcess(settings)

    runner.crawl(WeaponSpider)
    runner.start()


if __name__ == "__main__":
    main()
