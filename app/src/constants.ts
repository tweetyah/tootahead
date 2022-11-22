import type { TimeRange } from "./models";

export enum COMP_STATE {
  NONE = 0,
  LOADING = 1
}

export const TIME_RANGE: {[key: string]:TimeRange} = {
  // 12am - 6am
  NIGHT: { start: 0, end: 21600 },
  // 56am - 12pm
  MORNING: { start: 21600, end: 43200 },
  // 12pm - 6pm
  AFTERNOON: { start: 43200, end: 64800 },
  // 6pm - 12am
  EVENING: { start: 64800, end: 86399 }
}