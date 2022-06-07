from rox.server.rox_server import Rox, RoxOptions
from rox.server.flags.rox_flag import RoxFlag
from six import text_type
from typing import Dict, Optional, Text, Any

def main():
  api_key = get_api_key()

  flags = _FlagsContainer()

  Rox.register("", flags)
  
  Rox.setup(api_key).result()

  print('enableTutorial is {}'.format(flags.enableTutorial.is_enabled()))

  Rox.shutdown()

def get_api_key():
  text_file = open("api_key", "r")
  data = text_file.read()
  text_file.close()
  return data

class _FlagsContainer:
  def __init__(self):
    # type: () -> None
    # This is a special flag that defaults to False but we always set to True.
    # If its off it means we are not getting updated flag values from rollout
    self.enableTutorial = Flags(
        False, u"enableTutorial"
    )  # type: Flags

class Flags(RoxFlag):
  # def __init__(self):
  #   #Define the feature flags
  #   self.enableTutorial = RoxFlag(False)
  def __init__(self, default_value, name):
    # type: (bool, Text) -> None
    super(Flags, self).__init__(default_value)
    self.flag_name = name

  def is_enabled(self, context=None):
    # type: (Optional[Dict[Text, Any]]) -> bool
    _convert_context_int_to_text(context)
    flag_value = super(Flags, self).is_enabled(context)

    return flag_value

def _convert_context_int_to_text(context):
  # type: (Optional[Dict[Text, Any]] ) -> None
  if context is not None:
    for context_key in context:
      if isinstance(context[context_key], int):
        context[context_key] = text_type(context[context_key])

main()