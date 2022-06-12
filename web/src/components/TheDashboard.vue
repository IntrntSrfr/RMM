<template>
  <div class="dashboard">
    <div class="header">
      <fa-icon icon="bars" @click="openMenu" />
      <div>Servers</div>
    </div>
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
.header {
  display: flex;
  align-items: center;
  margin-bottom: 1em;
}

.header svg {
  height: 24px;
  width: 24px;
  padding: 0.5em;
}

.header div {
  font-size: 1.5rem;
  margin-left: 0.25em;
}
</style>
