<script>
import dataset from './dataset';

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
  {#each {length: 6} as _, i}
    <tr>
      <td class="time-index">{i + 1}</td>
      {#each ["monday", "tuesday", "wednesday", "thursday", "friday"] as day}
        <td on:click={() => openLink(dataset[day][i + 1]["link"])}>
          {dataset[day][i + 1]["name"]}
        </td>
      {/each}
    {/each}
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