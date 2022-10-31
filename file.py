from asyncio import run, sleep as async_sleep, to_thread as as_async
from os import environ as set_environment_variable

from pygame.event import get as get_joystick_events
import pyautogui as gui_automation, pydirectinput
from pygame.joystick import Joystick
from pygame import init


set_environment_variable["SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS"] = "1"


def get_key(keys: list[str], key: str) -> str:
    for (idx, k) in enumerate(keys):
        if k == key:
            return keys[idx]


def press_key(key: str) -> None:
    pydirectinput.press(key)


async def main(key: str):
    print("Capturing joystick events...")
    while True:
        for event in get_joystick_events():
            if hasattr(event, "button") and event.button == 10 and event.type == 1539:
                await as_async(press_key, key=key)
                break
        await async_sleep(0.016)


if __name__ == "__main__":
    key = get_key(keys=gui_automation.KEY_NAMES, key="insert")
    init()
    joystick = Joystick(0)
    joystick.init()
    try:
        run(main(key=key), debug=True)
    except KeyboardInterrupt:
        ...
