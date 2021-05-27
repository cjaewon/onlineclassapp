<script>
import { timetable } from "./store";

function locate(link) {
  window.open(link);
}

function maxLength() {
  if ($timetable) {

    return [
      $timetable["monday"].length, 
      $timetable["tuesday"].length, 
      $timetable["wednesday"].length, 
      $timetable["thursday"].length, 
      $timetable["friday"].length,
    ].reduce((prev, curr) => prev < curr ? curr : prev);
  } else {
    return 0;
  }
}
</script>

<table>
  <thead>
    <tr>
      <th></th>
      <th>월</th>
      <th>화</th>
      <th>수</th>
      <th>목</th>
      <th>금</th>
    </tr>
  </thead>
  <tbody>
    {#if $timetable}
      {#each {length: maxLength()} as _, i}
        <tr>
          <td class="time-index">{i + 1}</td>
          {#each ["monday", "tuesday", "wednesday", "thursday", "friday"] as day}
            {#if $timetable[day][i]}
              <td on:click={() => locate($timetable[day][i]["link"])}>
                {$timetable[day][i]["name"]}
              </td>
            {/if}
          {/each}
        {/each}
    {/if}
  </tbody>
</table>

<style>
  table {
    border: 1px solid black;
    border-radius: 8px;
    width: 500px;
    padding: 0.5rem 1rem;
}

  td {
    text-align: center;
    padding: 0.25rem 0;
    border: 1px solid black;
    margin: 0.5rem 0.5rem;
    width: 70px;
    cursor: pointer;
  }

  .time-index {
    width: 30px;
  }
</style>