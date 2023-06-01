<script lang="ts">
  import { onMount } from "svelte";
  import Item from "./Item.svelte";

  interface Shrt {
    id: number;
    redirect: string;
    shrt: string;
    clicked: number;
    random: boolean;
  }

  let shrts: Shrt[] = [];

  onMount(async () => {
    const res = await fetch("http://127.0.0.1:8080/shrts");
    shrts = await res.json();
  });
</script>

{#each shrts as shrt (shrt.id)}
  <Item {shrt} />
{/each}
