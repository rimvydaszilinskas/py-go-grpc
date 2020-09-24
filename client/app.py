from flask import Flask, jsonify
import grpc

import countries_pb2
import countries_pb2_grpc

app = Flask(__name__)

def obj_to_dict(obj): return obj.__dict__

@app.route("/<country>")
def country(country):
    print("Start service")
    try:
        channel = grpc.insecure_channel('localhost:3000')
        stub = countries_pb2_grpc.CountryStub(channel)
        countryRequest = countries_pb2.CountryRequest(name=country)
        countryResponse = stub.Search(countryRequest)
        return jsonify({
            "name": countryResponse.name,
            "alpha2Code": countryResponse.alpha2Code,
            "capital": countryResponse.capital,
            "subregion": countryResponse.subregion,
            "population": countryResponse.population,
            "nativeName": countryResponse.nativeName
        })
    except Exception as e:
        return jsonify({
            'error': 'error'
        })

if __name__ == '__main__':
    app.run()
