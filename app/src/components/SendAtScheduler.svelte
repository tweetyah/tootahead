<script lang="ts">
  import Button from "./Button.svelte";
  import { TIME_RANGE } from '../constants'
  import type { TimeRange } from "../models";
  import { onMount } from "svelte";

  export let value: Date
  let _date: string
  let _time: string

  let dates: Date[] = []

  onMount(() => {
    function addDays(date: Date, daysToAdd: number) {
      date.setDate(date.getDate() + daysToAdd);
      return date;
    }

    let d: Date[] = []
    for(let i = 0; i < 8; i++) {
      d.push(addDays(new Date(), i))
    }
    dates = d
  })

  function setDate(date: Date) {
    let dateStr = `${date.getFullYear()}-`
    dateStr += date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}-` : `${date.getMonth() + 1}-`
    dateStr += date.getDate() < 10 ? `0${date.getDate()}` : `${date.getDate()}`
    let timeStr = date.getHours() < 10 ? `0${date.getHours()}:` : `${date.getHours()}:`
    timeStr += date.getMinutes() < 10 ? `0${date.getMinutes()}` : `${date.getMinutes()}`
    value = date
    _date = dateStr
    _time = timeStr
  }

  function setTimeRange(range: TimeRange) {
    let min = range.start
    let max = range.end
    let seconds = Math.floor(Math.random() * (max - min + 1) + min)
    // Set it to the minute by removing the remaining seconds
    seconds = seconds - (seconds % 60)
    const newDate = new Date(
      value.getFullYear(),
      value.getMonth(),
      value.getDate(),
      0,
      0,
      0,
      0
    )
    newDate.setSeconds(seconds)
    setDate(newDate)
  }

  function setSelectedDate(date: Date) {
    let newDate = new Date(
      date.getFullYear(),
      date.getMonth(),
      date.getDate(),
      value.getHours(),
      value.getMinutes(),
      0,
      0
    )
    setDate(newDate)
  }

  let onDateInputChangedTimer
  function onDateInputChanged(e) {
    clearTimeout(onDateInputChangedTimer);
    onDateInputChangedTimer = setTimeout(() => {
      let spl = e.target.value.split("-")
      let d = new Date(spl[0], spl[1], spl[2], 0, 0, 0)
      setSelectedDate(d);
    }, 300);
  }

  let onTimeInputChangedTimer
  function onTimeInputChanged(e) {
    clearTimeout(onTimeInputChangedTimer);
    onTimeInputChangedTimer = setTimeout(() => {
      let spl = e.target.value.split(":")
      const newDate = new Date(
        value.getFullYear(),
        value.getMonth(),
        value.getDate(),
        spl[0],
        spl[1],
        0,
        0
      )
      setDate(newDate);
    }, 300);
  }
</script>

<div class="mb-2 p-2">
  <div class="border-[1px] border-mastodon rounded p-1 flex gap-4">
    <input class="focus:outline-0" type="date" bind:value={_date} on:change={onDateInputChanged} />
    <input class="focus:outline-0" type="time" bind:value={_time} on:change={onTimeInputChanged}/>
  </div>
  <h3>Date</h3>
  <div class="grid grid-cols-2 gap-2 mb-2">
    {#each dates as d}
      <Button variant="outlined" title={d.toLocaleDateString()} onClick={() => setSelectedDate(d)} />
    {/each}
  </div>
  <h3>Time</h3>
  <div class="grid grid-cols-2 gap-2">
    <Button variant="outlined" title="Night (12-6am)" onClick={() => setTimeRange(TIME_RANGE.NIGHT)} />
    <Button variant="outlined" title="Morning" onClick={() => setTimeRange(TIME_RANGE.MORNING)} />
    <Button variant="outlined" title="Afternoon" onClick={() => setTimeRange(TIME_RANGE.AFTERNOON)} />
    <Button variant="outlined" title="Evening" onClick={() => setTimeRange(TIME_RANGE.EVENING)} />
  </div>
</div>


