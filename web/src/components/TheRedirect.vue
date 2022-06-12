<template>
  <div>Redirecting...</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useUserStore } from "../stores/user";

export default defineComponent({
  name: "TheRedirect",
  setup() {
    const userStore = useUserStore();
    return { userStore };
  },
  async created() {
    const code = this.$route.query.code;
    if (!code) {
      this.$router.push("/");
      return;
    }

    try {
      await this.userStore.oauth(code.toString());
      await this.userStore.fetchUser();
    } catch (error) {
      this.$router.push("/");
      console.log(error);
    }
    this.$router.push("/dashboard");
  },
});
</script>

<style scoped></style>
