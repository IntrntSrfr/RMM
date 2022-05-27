<template>
  <div class="dashboard">
    <h1>Dashboard</h1>
    <guild-list />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "../stores/user";
import GuildList from "./GuildList.vue";

export default defineComponent({
  name: "TheDashboard",
  components: { GuildList },
  setup() {
    const userStore = useUserStore;

    return {};
  },
});
</script>

<style scoped></style>
