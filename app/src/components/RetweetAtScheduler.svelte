<script lang="ts">
  import Button from "./Button.svelte";
  import { TIME_RANGE } from '../constants'
  import type { TimeRange } from "../models";
  import { onMount } from "svelte";

  export let value: Date
  export let isEnabled: boolean = false

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

    setDate(new Date())
  })

  function setDate(date: Date) {
    let dateStr = `${date.getFullYear()}-`
    dateStr += date.getMonth() < 10 ? `0${date.getMonth()}-` : `${date.getMonth()}-`
    dateStr += date.getDay() < 10 ? `0${date.getDate()}` : `${date.getDay()}-`
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

</script>

<div class="mb-2 p-2">
  <div>
    <div class="relative inline-block w-10 mr-2 align-middle select-none transition duration-200 ease-in">
      <input bind:checked={isEnabled} type="checkbox" name="toggle" id="toggle" class="toggle-checkbox absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer"/>
      <label for="toggle" class="toggle-label block overflow-hidden h-6 rounded-full bg-slate-300 cursor-pointer"></label>
    </div>
    <label for="toggle">Enable retweeting</label>
  </div>
  {#if isEnabled}
    <div class="border-[1px] border-slate-300 rounded p-1 flex gap-4 mt-2 mb-1">
      <input class="focus:outline-0" type="date" bind:value={_date} />
      <input class="focus:outline-0" type="time" bind:value={_time}/>
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
  {/if}
</div>


