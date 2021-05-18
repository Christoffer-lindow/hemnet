<template>
  <div class="home">
    <div class="sidebar">
        <div class="sidebar-btn" :class="activeLink == 'overview' && 'active'" @click="setActive('overview')">Overview</div>
        <div class="sidebar-btn" :class="activeLink == 'gavle_dataset' && 'active'" @click="setActive('gavle_dataset')">Gävle Dataset</div>
        <div class="sidebar-btn" :class="activeLink == 'stockholm_dataset' && 'active'" @click="setActive('stockholm_dataset')">Stockholm Dataset</div>
        <div class="sidebar-btn" :class="activeLink == 'property_description' ? 'active' : 'sidebar-btn-disabled'">Property Description</div>
    </div>
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
export default {
  name: "Home",
  components: {
    PropertyTable,
    PropertyDescription,
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
      }
  }
};
</script>
<style  scoped>
/* The side navigation menu */
.sidebar {
  margin: 0;
  padding: 0;
  width: 225px;
  background-color: white;
  position: fixed;
  height: 100%;
  overflow: auto;
}
.sidebar-btn {
    width: 100%;
    padding-top: 10px;
    padding-bottom: 10px;
    cursor: pointer;
}
.active {
    background-color: green;
    color: white;
}
/* Page content. The value of the margin-left property should match the value of the sidebar's width property */
div.content {
  margin-left: 225px;
  height: 1000px;
}

/* On screens that are less than 700px wide, make the sidebar into a topbar */
@media screen and (max-width: 700px) {
  .sidebar {
    width: 100%;
    height: auto;
    position: relative;
  }
  .sidebar-btn {
    float: left;
  }
  div.content {
    margin-left: 0;
  }
}

/* On screens that are less than 400px, display the bar vertically, instead of horizontally */
@media screen and (max-width: 400px) {
  .sidebar-btn {
    text-align: center;
    float: none;
  }
}
.sidebar-btn-disabled {
    cursor: default;
    color: #b9b9b9
}
</style>
