<template>
  <div class="dashboard">
    <h1><fa-icon icon="bars" @click="openMenu" /> Servers</h1>
    <guild-list />
  </div>
</template>

<script lang="ts">
import { useNavStore } from "@/stores/navbar";
import { defineComponent } from "vue";
import GuildList from "./GuildList.vue";

export default defineComponent({
  name: "TheDashboard",
  components: { GuildList },
  setup() {
    const navStore = useNavStore();
    return { navStore };
  },
  methods: {
    openMenu() {
      this.navStore.setActive(true);
    },
  },
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
