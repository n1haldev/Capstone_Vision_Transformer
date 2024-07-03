import sys
from scrape_google_images import scrape_google_images

def scraper(term):
    scrape_google_images(
        download_path="./Google Images/",
        terms=[term],
        number_of_images=300,
        resolution=[256, 256],
        log=True
    )

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python scraper.py <plant_name>")
        sys.exit(1)

    plant_name = sys.argv[1]
    scraper(plant_name)

