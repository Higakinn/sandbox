#ライブラリをインポート
from selenium import webdriver
from selenium.webdriver.common.by import By

import time

options = webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--no-sandbox')
options.add_argument('--disable-dev-shm-usage')

keyword="test"
print("\nブラウザを設定")
driver = webdriver.Chrome('/usr/local/bin/chromedriver',options=options)
driver.implicitly_wait(10)

driver.get(f"https://www.google.com/search?q=新居浜")

# 地図のタブに移動する
map_field='/html/body/div[7]/div/div[4]/div/div[1]/div/div[1]/div/div[2]/a'
elem = driver.find_element(By.XPATH,map_field)

elem.click()
time.sleep(1)
driver.save_screenshot(f"{keyword}.png")
print(driver.current_url)
