<template>
  <div class="member-wrapper">
    <div
      class="btn scroll-top"
      :class="{ active: showScroll }"
      @click="scrollUp()"
    >
      <fa-icon icon="arrow-up" />
    </div>
    <p class="counter">
      {{ selectedMembers.length }} / {{ members.length }} selected
    </p>
    <!-- <input type="text" v-model="search" @input="" /> -->
    <div class="actions">
      <div class="grp">
        <AppButton text="Select none" @click="markNone" />
        <AppButton text="Select all" @click="markAll" />
      </div>
      <div class="grp">
        <AppButton text="Ban selected" variant="danger" @click="banSelected" />
      </div>
    </div>
    <AppLoader v-if="loadingMembers" />
    <div v-else class="member-list">
      <GuildMembersItem
        v-for="(m, i) in members"
        :key="i"
        @click="mark(m)"
        :checked="m.checked"
        :id="m.member.user.id"
        :username="`${m.member.user.username}#${m.member.user.discriminator}`"
        :icon="m.member.user.avatar"
        :joined="m.joined"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { Member } from "@/stores/user";
import AppButton from "@/components/AppButton.vue";
import { useGuildStore } from "@/stores/guilds";
import GuildMembersItem from "./GuildMembersItem.vue";
import AppLoader from "./AppLoader.vue";

type TableMember = {
  member: Member;
  checked: boolean;
  joined: Date;
};

export default defineComponent({
  name: "GuildMembers",
  components: { AppButton, GuildMembersItem, AppLoader },
  setup() {
    const guildStore = useGuildStore();
    return { guildStore };
  },
  data() {
    return {
      showScroll: false,
      search: "",
      members: [] as TableMember[],
    };
  },
  computed: {
    selectedMembers(): TableMember[] {
      return this.members.filter((m: TableMember) => m.checked);
    },
    loadingMembers(): boolean {
      const guildID = this.$route.params.guildID;
      const g = this.guildStore.getGuildByID(guildID.toString());
      if (!g) {
        return true;
      }
      return !g.members;
    },
  },
  methods: {
    markAll() {
      this.members.forEach((m: TableMember) => (m.checked = true));
    },
    markNone(): void {
      this.members.forEach((m: TableMember) => (m.checked = false));
    },
    mark(tm: TableMember) {
      tm.checked = !tm.checked;
    },
    banSelected() {
      if (!this.selectedMembers.length) {
        return;
      }
      const ans = confirm("Are you SURE you want to ban ALL selected users?");
      if (!ans) {
        return;
      }

      console.log(this.selectedMembers.map((m) => m.member.user.id).join(" "));
    },
    handleScroll() {
      this.showScroll = window.scrollY > 300;
    },
    scrollUp(){
      window.scrollTo({ top: 0, behavior: "smooth" });
    },
  },
  async created() {
    const guildID = this.$route.params.guildID;
    const g = this.guildStore.getGuildByID(guildID.toString());
    await this.guildStore.fetchGuildMembers(guildID.toString());

    if (!g || !g.members) {
      return;
    }

    this.members = g.members
      .map((m: Member) => {
        return {
          member: m,
          checked: false,
          joined: new Date(m.joined_at),
        } as TableMember;
      })
      .sort((a: TableMember, b: TableMember) => {
        return b.joined.getTime() - a.joined.getTime();
      });
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    window.removeEventListener("scroll", this.handleScroll);
  },
});
</script>

<style scoped>
.counter {
  margin-bottom: 0.5em;
}

.actions {
  display: flex;
  justify-content: space-between;
  gap: 1em;
  padding-bottom: 1em;
  border-bottom: 1px solid #666;
}

.grp {
  display: flex;
  gap: 1em;
}

.member-list {
  margin-top: 0.5em;
  display: flex;
  flex-direction: column;
  width: 100%;
}

.member + .member {
  border-top: 1px solid dodgerblue;
}

.scroll-top {
  cursor: pointer;
  display: flex;
  position: fixed;
  bottom: -1em;
  right: 2em;

  background-color: rgb(74, 60, 92);

  padding: 0.5em;
  border-radius: 50%;

  visibility: hidden;
  opacity: 0;
  z-index: 9999;
  transition: 0.2s;
}

.scroll-top:hover {
  background-color: rgb(119, 70, 184);
}

.scroll-top.active {
  visibility: visible;
  opacity: 1;
  bottom: 2em;
}

.scroll-top svg {
  height: 36px;
  width: 36px;
}
</style>
