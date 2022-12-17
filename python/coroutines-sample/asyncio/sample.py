import asyncio
from contextlib import asynccontextmanager
from typing import Any, AsyncGenerator, AsyncIterator, Dict, Optional
from urllib.parse import urlparse

class Broadcast:
    def __init__(self):
        self._backend = None
        self._subscribers: Dict[str, Any] = {}

    async def __aenter__(self) -> "Broadcast":
        # await self.connect()
        print("__aenter__")
        return self

    async def __aexit__(self, *args: Any, **kwargs: Any) -> None:
        print("__aexit__")
        # await self.disconnect()

    async def connect(self) -> None:
        # await self._backend.connect()
        print("connect")
        self._listener_task = asyncio.create_task(self._listener())

    async def disconnect(self) -> None:
        if self._listener_task.done():
            self._listener_task.result()
        else:
            self._listener_task.cancel()
        await self._backend.disconnect()

    async def _listener(self) -> None:
        while True:
            # event = await self._backend.next_published()
            event = "test"
            for queue in list(self._subscribers.get(event.channel, [])):
                await queue.put(event)

    async def publish(self, channel: str, message: Any) -> None:
        await self._backend.publish(channel, message)

    @asynccontextmanager
    async def subscribe(self, channel: str) -> AsyncIterator["Subscriber"]:
        queue: asyncio.Queue = asyncio.Queue()
        print(queue)
        try:
            # if not self._subscribers.get(channel):
            #     print("test")
            #     await self._backend.subscribe(channel)
            #     self._subscribers[channel] = set([queue])
            # else:
            #     self._subscribers[channel].add(queue)

            yield Subscriber(queue)

            # self._subscribers[channel].remove(queue)
            # if not self._subscribers.get(channel):
            #     del self._subscribers[channel]
            #     await self._backend.unsubscribe(channel)
        finally:
            await queue.put(None)

class Subscriber:
    def __init__(self, queue: asyncio.Queue) -> None:
        self._queue = queue

    async def __aiter__(self) -> Optional[AsyncGenerator]:
        try:
            while True:
                print("待機中")
                yield await self.get()
        except Exception:
            pass
    # Event
    async def get(self) -> Any:
        item = await self._queue.get()
        if item is None:
            raise Exception()
        return item

broadcast = Broadcast()

async def main():
  async with Broadcast() as s:
    print(s)
  async with broadcast.subscribe(channel="chatroom") as subsciber:
    async for event in subsciber:
      # yield event.message
      print(event)
  

  # queue = asyncio.Queue()
  # sub = Subscriber(queue)


# entrypointRP
asyncio.run(main())


