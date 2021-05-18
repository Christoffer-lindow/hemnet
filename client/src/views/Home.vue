<template>
  <div class="home">
      <side-bar v-on:sidebarClicked="handleSidebarClicked" :active-link="activeLink" />
    <div class="content">
      <property-table :city="'gavle'" v-if="activeLink == 'gavle_dataset'" v-on:tableClicked="handleTableClicked" />
      <property-table :city="'stockholm'" v-if="activeLink == 'stockholm_dataset'" v-on:tableClicked="handleTableClicked" />
      <property-description :property="property" v-if="activeLink == 'property_description'" />
    </div>
  </div>
</template>

<script>
import PropertyTable from "../components/PropertyTable.vue";
import PropertyDescription from "../components/PropertyDescription.vue"
import SideBar from "../components/Sidebar.vue"
export default {
  name: "Home",
  components: {
    PropertyTable,
    PropertyDescription,
    SideBar
  },
  data() {
      return {
          activeLink: "overview",
          property: null
      }
  },
  methods: {
      setActive(component) {
          this.activeLink = component
      },
      handleTableClicked(property) {
          this.property = property
          this.setActive("property_description")
      },
      handleSidebarClicked(component) {
          this.setActive(component)
      }
  }
};
</script>
<style  scoped>

</style>
