<template>
  <div
    class="btn"
    :class="{ outlined: variant === 'outlined', danger: variant === 'danger' }"
  >
    {{ text }}
  </div>
</template>

<script setup lang="ts">
import { defineProps, withDefaults } from "vue";

export interface Props {
  text: string;
  variant?: string;
}

withDefaults(defineProps<Props>(), { variant: "outlined" });
</script>

<style scoped>
.btn {
  position: relative;
  cursor: pointer;
  border: none;
  font-size: 0.8rem;
  font-family: Inter, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    Oxygen, Ubuntu, Cantarell, "Fira Sans", "Droid Sans", "Helvetica Neue",
    sans-serif;
  padding: 0.5em 0.75em;
  border-radius: 5px;
  transition: 0.2s;
}

.outlined {
  --text-color: white;
  --border-color: dodgerblue;
  --border-hover-color: rgba(22, 115, 213, 0.2);
  color: var(--text-color);
}

.outlined:after {
  content: "";
  position: absolute;
  height: 100%;
  width: 100%;
  left: 0;
  top: 0;
  border: 1px solid var(--border-color);
  border-radius: 5px;
}

.outlined:hover {
  background-color: var(--border-hover-color);
}

.danger {
  background-color: rgb(232, 29, 29);
}

.danger:hover,
.danger:focus {
  background-color: rgb(207, 25, 25);
}
</style>
