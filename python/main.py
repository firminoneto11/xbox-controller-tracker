from asyncio import run, sleep as async_sleep, to_thread as as_async
from os import environ as set_environment_variable

from pygame.event import get as get_joystick_events
from pygame.joystick import Joystick
from pygame import init
import pydirectinput


set_environment_variable["SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS"] = "1"


def press_key(key: str) -> None:
    pydirectinput.press(key)


async def main(key: str):
    print("Capturing joystick events...")
    while True:
        for event in get_joystick_events():
            if hasattr(event, "button") and event.button == 10 and event.type == 1539:
                await as_async(press_key, key=key)
                print(f"'{key}' was pressed!")
                break
        await async_sleep(0.010)


if __name__ == "__main__":
    key = "insert"
    init()
    joystick = Joystick(0)
    joystick.init()
    try:
        run(main(key=key), debug=True)
    except KeyboardInterrupt:
        ...
