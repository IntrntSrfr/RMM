<template>
  <div class="guild">
    <div class="name">
      {{ name }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "GuildListItem",
  props: {
    name: { type: String, required: true },
    icon: { type: String, default: "" },
  },
});
</script>

<style scoped>
.guild {
  padding: 1em;
}

@media only screen and (min-width: 660px) {
  .guild {
    border: 1px solid dodgerblue;
    border-radius: 5px;
  }
}
</style>
