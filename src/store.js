import { writable } from "svelte/store";

export const timetable = writable(JSON.parse(localStorage.getItem("timetable")));

timetable.subscribe((table) => {
  localStorage.setItem("timetable", JSON.stringify(table));
});