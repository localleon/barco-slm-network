# This file is used primary for configuration purpose
users = {
    "FMT": "test",                                          # FrontEnd Default User
}

sessionlog = []  # Array for logging entrys

success_response = {
    "status": "success",
}

serial_commands = {
    "shutter_on": "turn on shutter cmd",
    "shutter_off": "turn off shutter cmd",
}

# NOTE: These numbers are 0-based because we are handling 0-based arrays/lists
sacn_mapping = {
    0: "shutter_on",
    1: "shutter_off",
}


def dataresponse(data):
    return{
        "data": data,
    }
