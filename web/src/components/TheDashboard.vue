<template>
  <div class="dashboard">
    <h1><fa-icon icon="bars" /> Servers</h1>
    <guild-list />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GuildList from "./GuildList.vue";

export default defineComponent({
  name: "TheDashboard",
  components: { GuildList },
});
</script>

<style scoped>
h1 {
  margin-bottom: 0.5em;
}

h1 svg {
  margin-right: 0.5em;
}
</style>
