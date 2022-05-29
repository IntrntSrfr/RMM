<template>
  <div class="dashboard">
    <h1>Dashboard</h1>
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

<style scoped></style>
