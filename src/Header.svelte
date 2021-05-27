<script>
import { timetable } from './store';

function editTimeTable() {
  const input = prompt();

  try {
    const parsed = JSON.parse(input);
    timetable.update(() => parsed);
  } catch (e) {
    alert("포맷에 맞게 시간표를 입력해주세요.");
  }
}
</script>

<header>
  <button on:click={() => window.open($timetable["homeroom-teacher"])}>담임T</button>
  <button on:click={editTimeTable}>시간표 변경</button>
</header>

<style>
  header {
    margin: 0.5rem 0;
    display: flex;
    justify-content: space-between;
    width: 500px;
  }

  button {
    border: 1px solid black;
    border-radius: 4px;
    padding: 0.45rem 0.45rem;
    background: transparent;
    cursor: pointer;
  }
</style>