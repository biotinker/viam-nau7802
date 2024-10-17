# viam-nau7802

A Viam module to use a nau7802 load cell amplifier

## Setup

Add the module to your robot from the Viam registry. The component may then be added to your configuration as follows (using example values)

```
  ...
  "components": [
    ...,
    {
      "name": "componentNameHere",
      "model": "biotinker:sensor:nau7802",
      "type": "sensor",
      "namespace": "biotinker",
      "attributes": {
        "zero": 161000,
        "kg": 50000,
        "gain": 128,
        "samples": 16,
        "i2c_bus": "0"
      }
    },
    ...,
  ],
  ...
```

The attributes `gain` and `samples` are optional.

If `gain` is not set, it will default to 64. Valid `gain` values are 1, 2, 4, 8, 16, 32, 64, and 128. Gain should be set higher for lower capacity load cells. Any change to `gain` requires re-calibration.

If `samples` is not set, it will default to 8. This value is the number of readings which will be taken in quick succession and averaged together to give the reading. Valid `samples` values is any positive integer.

## Calibration

The output of the sensor has two values, a raw reading and a kg reading. The kg reading will be meaningless without proper calibration.

First, assemble your apparatus. Take a reading with nothing on the scale. This raw value should be entered into the `zero` field in your configuration.

Then, place an item of known weight on the scale and record the raw value. The value to fill in for `kg` is the amount that the raw value is expected to change for each kg of weight. For example, if your zero value is 163540, and you place a 200g item on the scale and that raw value is 167315, then your kg value should be `(167315-163540) * (1000g/200g)` which equals `18875`

The `kg` value may be negative, for example if your load cell is installed upside down.
