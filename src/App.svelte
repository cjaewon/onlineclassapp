<script>
import Header from "./Header.svelte";
import TimeTable from "./TimeTable.svelte";

</script>

<main>
	<div>
		<h2>학급 시간표</h2>
		<Header />
		<TimeTable />
	</div>
</main>

<style>
	main {
		display: grid;
		place-items: center;
		height: 80vh;
	}
</style>
