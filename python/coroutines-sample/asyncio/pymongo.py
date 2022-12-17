import asyncio
import concurrent.futures
from contextlib import closing
import logging

import pymongo

logging.basicConfig(level=logging.DEBUG,
                    format='%(asctime)s %(process)d(%(thread)s) [%(levelname)s] %(message)s')


def sleep(delay):
    logging.info(f'{delay}-start')
    with closing(mysql.connect()) as dbconn:
        with dbconn.cursor() as cursor:
            cursor.execute(f'SELECT SLEEP({delay})')
    logging.info(f'{delay}-end')


async def sleep_async(delay):
    return sleep(delay)


loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.gather(sleep_async(2), sleep_async(4), sleep_async(6)))