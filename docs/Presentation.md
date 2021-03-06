---
theme : "night"
transition: "convex"
title: "Deduplication"
highlightTheme: "darkula"
logoImg: "./redis-labs-header.png"
slideNumber: false
---

## De-Duplication

---

## Agenda

<small>

* Use Cases
* Why Redis?
* Why use a module?
* Bloom Filters
* Cuckoo Filters
* Demo


</small>

---

## Use Cases

- Messaging
- Update Limiting
- ETL Processing
- Ad Serving

---

### Messaging
<br>
<small>

- Large online tools company in Fortune 50
- Sends updates and notifications to subscribers
- Customers unsubscribe when receiving duplicate updates
- Each message is a Bloom Filter with customer id as a member

</small>

---

### Update Limiting
<br>
<small>

- Set timed bucket with at TTL 
- Useful to combat alert fatigue
- Often around monitoring
- Critical_DBAlert_2020-01-06-1300

</small>

---

### ETL Processing
<br>
<small>

- asynchronous processing of logs or events
- Each duplicate costs more in processing and storage
- Usually large fitlers over time buckets

</small>

---

### Ad Serving
<br>
<small>

- Ad Networks often get paid for unique impressions
- Can be coupled with other modules to maximize revenue
- Main user focuses on Cuckoo filters
- [AdGears Demo](https://github.com/Redislabs-Solution-Architects/AdGears)

</small>

---

## Why Redis?
<br>
<small>

- centralized filter as opposed to several local ones
- inherits all of the HA, observability and automation of Redis
- Redis runs at ludicrous speed

</small>

---

## Why Modules? 
### Speed!

<br>

<small>

|Implementation|Add|Check|
|--|--|--|--|
|redablooms | 20k/s | 7k/s|
|lua | 29k/s | 25k/s|
|bloomd | 250k/s | 200k/s|
|redisbloom | 400k/s | 400k/s|
|  |  |  |

</small>

---

## Why Modules? 
### Efficiency!

<br>
<small>

| Benchmark | RedisBloom | Set |
|--|--|--|--|
|Size | 599,126 bytes | 18,000,005 bytes|
| 1M Checks | 1m23.81s | 1m26.85s |
|  |  | | |

</small>
<br>

---

## Bloom Filters

- Probabilistic data structure
- Trade accuracy for speed/efficiency
- Answer of "no" is 100% accurate
- Answer of "yes" has an error associated with it

---

## Cuckoo Filters

- Probabilistic data structure
- Trade accuracy for speed/efficiency
- Answer of "no" is 100% accurate
- Answer of "yes" has an error associated with it

---

# Wait a minute!!

---

## Bloom Filter

<img src="./bloom-1.png" style="background:none; border:none; box-shadow:none;">

---

## Bloom Filter

<img src="./bloom-2.png" style="background:none; border:none; box-shadow:none;">

---

## Bloom Filter

<img src="./bloom-3.png" style="background:none; border:none; box-shadow:none;">

---

## Bloom Filter

<img src="./bloom-4.png" style="background:none; border:none; box-shadow:none;">

---

## Bloom Filter

<img src="./bloom-5.png" style="background:none; border:none; box-shadow:none;">

---

## Cuckoo Filter

<img src="./cuckoo-1.png" style="background:none; border:none; box-shadow:none;">

---

## Cuckoo Filter

<img src="./cuckoo-2.png" style="background:none; border:none; box-shadow:none;">

---

## Cuckoo Filter

<img src="./cuckoo-3.png" style="background:none; border:none; box-shadow:none;">

---

## Cuckoo Filter

<img src="./cuckoo-4.png" style="background:none; border:none; box-shadow:none;">

---

## Cuckoo Filter

<img src="./cuckoo-5.png" style="background:none; border:none; box-shadow:none;">

---

## Cuckoo Filter

<img src="./cuckoo-6.png" style="background:none; border:none; box-shadow:none;">

---

## When to Use

<small>

- Cuckoo Filter
    - Need to delete data from the filter
    - Faster Checks
    - CF.ADDNX is slower
- Bloom Filter
    - Insert is cheaper and faster
    - Smaller size
    - Insert/Check single operation
    - No deletions necessary

</small>
    

---

# Demo

---

<img src="./webform.png" style="background:none; border:none; box-shadow:none;">

---

<img src="./demo-1.png" style="background:none; border:none; box-shadow:none;">

---

<img src="./demo-2.png" style="background:none; border:none; box-shadow:none;">

---

<img src="./demo-3.png" style="background:none; border:none; box-shadow:none;">

---

### The Gear

<pre><code>
def runIt(x):
    val = x['value']['Value']
    name = x['value']['Name']
    execute('TS.INCRBY', 's-unfiltered', 1, 'TIMESTAMP', '*')
    j = execute("BF.ADD", "DEDUP", val)
    if j > 0:
        execute('TS.INCRBY', 's-filtered', 1, 'TIMESTAMP', '*')

gb =  GearsBuilder(
        reader = 'StreamReader',
        desc   = "Process messages")

gb.map(runIt)
gb.register('MessageQueue')

</code></pre>

---

<img src="./demo-4.png" style="background:none; border:none; box-shadow:none;">

---

<img src="./demo-diagram.png" style="background:none; border:none; box-shadow:none;">

---

# Thanks