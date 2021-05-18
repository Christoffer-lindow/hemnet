<template>
  <div class="table">
    <div class="row header">
      <div class="cell" @click="orderBy('street')">
        <span class="row-cell-header"
          >Street
          <i class="fas fa-sort-up" v-if="lastFiltered !== 'street'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'street'"></i
        ></span>
      </div>
      <div class="cell" @click="orderBy('address')">
        <span class="row-cell-header"
          >Address
          <i class="fas fa-sort-up" v-if="lastFiltered !== 'address'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'address'"></i
        ></span>
      </div>
      <div class="cell" @click="orderBy('price')">
        <span class="row-cell-header"
          >Price sek <i class="fas fa-sort-up" v-if="lastFiltered !== 'price'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'price'"></i
        ></span>
      </div>
      <div class="cell" @click="orderBy('area')">
        <span class="row-cell-header"
          >Area <i class="fas fa-sort-up" v-if="lastFiltered !== 'area'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'area'"></i
        ></span>
      </div>
      <div class="cell" @click="orderBy('price_sqm')">
        <span class="row-cell-header"
          >Price/m²
          <i class="fas fa-sort-up" v-if="lastFiltered !== 'price_sqm'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'price_sqm'"></i
        ></span>
      </div>
      <div class="cell" @click="orderBy('rooms')">
        <span class="row-cell-header"
          >Rooms <i class="fas fa-sort-up" v-if="lastFiltered !== 'rooms'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'rooms'"></i
        ></span>
      </div>
      <div class="cell" @click="orderBy('rent')">
        <span class="row-cell-header"
          >Rent <i class="fas fa-sort-up" v-if="lastFiltered !== 'rent'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'rent'"></i
        ></span>
      </div>
      <div class="cell" @click="orderBy('year_built')">
        <span class="row-cell-header"
          >Build Year
          <i class="fas fa-sort-up" v-if="lastFiltered !== 'year_built'"></i
          ><i class="fas fa-sort-down" v-if="lastFiltered === 'year_built'"></i
        ></span>
      </div>
    </div>

    <div
      class="row"
      v-for="property in properties"
      :key="property.url"
      @click="emitEvent(property)"
    >
      <div class="cell" data-title="Street">{{ property.street }}</div>
      <div class="cell" data-title="Address">{{ property.address }}</div>
      <div class="cell" data-title="Price">{{ property.price }}</div>
      <div class="cell" data-title="Area">{{ property.area }} m²</div>
      <div class="cell" data-title="Price/m²">{{ property.price_sqm }}</div>
      <div class="cell" data-title="Rooms">{{ property.rooms }}</div>
      <div class="cell" data-title="Rent">
        {{ property.rent !== null ? property.rent : "Unspecified" }}
      </div>
      <div class="cell" data-title="Build year">
        {{
          property.year_built !== null
            ? property.year_built.split("T")[0]
            : "Unspecified"
        }}
      </div>
    </div>
  </div>
</template>

<script>
import { getProperties } from "../requests/requests";
export default {
  name: "PropertyTable",
  props: ["city"],
  data() {
    return {
      properties: null,
      lastFiltered: "",
    };
  },
  methods: {
    orderBy(filter) {
      if (this.lastFiltered !== filter) {
        this.properties = this.properties.sort((a, b) =>
          a[filter] < b[filter] ? 1 : b[filter] < a[filter] ? -1 : 0
        );
        this.lastFiltered = filter;
      } else {
        this.properties = this.properties.sort((a, b) =>
          a[filter] > b[filter] ? 1 : b[filter] > a[filter] ? -1 : 0
        );
        this.lastFiltered = "";
      }
    },
    emitEvent(property) {
        this.$emit("tableClicked", property)
    }
  },
  async mounted() {
    const p = await getProperties(this.city);
    this.properties = p.data;
  },
};
</script>

<style lang="css" scoped>
a {
  width: 100%;
  text-decoration: none;
  color: inherit;
}
a:visited {
  color: inherit;
}

.table {
  margin: 0 0 40px 0;
  width: 100%;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
  display: table;
}
@media screen and (max-width: 1210px) {
  .table {
    display: block;
  }
}
.row {
  display: table-row;
  background: #f6f6f6;
}
.row:nth-of-type(odd) {
  background: #e9e9e9;
}
.row.header {
  font-weight: 900;
  color: #ffffff;
  background: #00227e;
}
@media screen and (max-width: 1210px) {
  .row {
    padding: 14px 0 7px;
    display: block;
  }
  .row > .header {
    padding: 0;
    height: 6px;
  }
  .row.header > .cell {
    display: none;
  }
  .row.cell {
    margin-bottom: 10px;
  }
  .row > .cell:before {
    margin-bottom: 5px;
    content: attr(data-title);
    min-width: 98px;
    font-size: 12px;
    line-height: 10px;
    font-weight: bold;
    text-transform: uppercase;
    color: #727171;
    display: block;
  }
}
.cell {
  padding: 6px 12px;
  display: table-cell;
  cursor: pointer;
}
@media screen and (max-width: 1370px) {
  .cell {
    padding: 2px 16px;
    display: block;
  }
}
.row-cell-header {
  display: flex;
}
.fas {
  margin-left: 10px;
}
.fa-sort-up {
  margin-top: 4px;
}
</style>