# This file is used primary for configuration purpose
users = {
    "FMT": "test",                                          # FrontEnd Default User
}

success_response = {
    "status": "success",
}

serial_commands = {
    "shutter_off":  b'\xfe\x01\x22\x42\x00\x65\xff',
    "shutter_on": b'\xfe\x01\x23\x42\x00\x66\xff',
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
