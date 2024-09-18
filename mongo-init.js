db.createUser({
    user: "mongouser",
    pwd: "mongopass",
    roles: [{
        role: "readWrite",
        db: "telemetry"
    }]
});

db = db.getSiblingDB('telemetry');

db.createCollection("telemetry_devices");

db.telemetry_devices.insertOne({
    device_id: "uuid11",
    status: "active"
});