"use strict";
// this file was automatically generated, DO NOT EDIT
Object.defineProperty(exports, "__esModule", { value: true });
exports.ToObject = exports.FromArray = exports.ParseNumber = exports.ParseDate = void 0;
// helpers
const maxUnixTSInSeconds = 9999999999;
function ParseDate(d) {
    if (d instanceof Date)
        return d;
    if (typeof d === "number") {
        if (d > maxUnixTSInSeconds)
            return new Date(d);
        return new Date(d * 1000); // go ts
    }
    return new Date(d);
}
exports.ParseDate = ParseDate;
function ParseNumber(v, isInt = false) {
    if (!v)
        return 0;
    if (typeof v === "number")
        return v;
    return (isInt ? parseInt(v) : parseFloat(v)) || 0;
}
exports.ParseNumber = ParseNumber;
function FromArray(Ctor, data, def = null) {
    if (!data || !Object.keys(data).length)
        return def;
    const d = Array.isArray(data) ? data : [data];
    return d.map((v) => new Ctor(v));
}
exports.FromArray = FromArray;
function ToObject(o, typeOrCfg = {}, child = false) {
    if (o == null)
        return null;
    if (typeof o.toObject === "function" && child)
        return o.toObject();
    switch (typeof o) {
        case "string":
            return typeOrCfg === "number" ? ParseNumber(o) : o;
        case "boolean":
        case "number":
            return o;
    }
    if (o instanceof Date) {
        return typeOrCfg === "string"
            ? o.toISOString()
            : Math.floor(o.getTime() / 1000);
    }
    if (Array.isArray(o))
        return o.map((v) => ToObject(v, typeOrCfg, true));
    const d = {};
    for (const k of Object.keys(o)) {
        const v = o[k];
        if (v === undefined)
            continue;
        if (v === null)
            continue;
        d[k] = ToObject(v, typeOrCfg[k] || {}, true);
    }
    return d;
}
exports.ToObject = ToObject;
//# sourceMappingURL=blob.js.map