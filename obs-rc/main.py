from fastapi import FastAPI, HTTPException
import uvicorn
from obswebsocket import obsws, requests
import requests as reqs

ws = obsws("ip", 4455, "password") # OBS Websocket
ws.connect()
app = FastAPI()

@app.get("/obs_rc")
async def root(level: str | None = None):
    match level:
        case "1":
            try:
                # Discord, messages, speedrunning site automation requests here...
                ws.call(requests.SetCurrentProgramScene(sceneName=level))
            except:
                raise HTTPException(status_code=500, detail="Internal Server Error")
            return True
        case "2":
            try:
                # Discord, messages, speedrunning site automation requests here...
                ws.call(requests.SetCurrentProgramScene(sceneName=level))
            except:
                raise HTTPException(status_code=500, detail="Internal Server Error")
            return True
        case "3":
            try:
                # Discord, messages, speedrunning site automation requests here...
                ws.call(requests.SetCurrentProgramScene(sceneName=level))
            except:
                raise HTTPException(status_code=500, detail="Internal Server Error")
            return True
        case "4":
            try:
                # Discord, messages, speedrunning site automation requests here...
                ws.call(requests.SetCurrentProgramScene(sceneName=level))
            except:
                raise HTTPException(status_code=500, detail="Internal Server Error")
            return True
        case "5":
            try:
                # Discord, messages, speedrunning site automation requests here...
                ws.call(requests.SetCurrentProgramScene(sceneName=level))
            except:
                raise HTTPException(status_code=500, detail="Internal Server Error")
            return True
        case _:
            return False

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=40000) # HTTP host
    ws.disconnect()