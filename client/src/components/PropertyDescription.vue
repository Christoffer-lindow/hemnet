<template>
  <div>
    <div class="container">
      <div class="img-container">
        <img :src="property.img_url" alt="" />
      </div>
      <div class="info-container">
        <div class="info-street">
          <h2>{{ property.street }}</h2>
          <small>{{ property.address }}</small>
          <h3>{{ property.price }} kr</h3>
          <div class="info-link">
            <a :href="property.url">Check out property on hemnet</a>
          </div>
        </div>
        <div class="info-property">
          <div class="info-pair">
            <div class="info-pair-key">Property type</div>
            <div class="info-pair-val">{{ property.type }}</div>
          </div>
          <div class="info-pair">
            <div class="info-pair-key">Contract type</div>
            <div class="info-pair-val">{{ property.contract_type }}</div>
          </div>
          <div class="info-pair">
            <div class="info-pair-key">Rooms</div>
            <div class="info-pair-val">{{ property.rooms }}</div>
          </div>
          <div class="info-pair">
            <div class="info-pair-key">Living area</div>
            <div class="info-pair-val">{{ property.area }} m²</div>
          </div>
          <div class="info-pair">
            <div class="info-pair-key">Balcony</div>
            <div class="info-pair-val">
              {{ property.balcony === true ? "Yes" : "No" }}
            </div>
          </div>
          <div class="info-pair">
            <div class="info-pair-key">Build year</div>
            <div class="info-pair-val">
              {{
                property.year_built !== null
                  ? property.year_built.split("T")[0]
                  : "Unspecified"
              }}
            </div>
          </div>
        </div>
        <div class="info-property">
          <div class="info-pair">
            <div class="info-pair-key">Rent</div>
            <div class="info-pair-val">{{ property.rent }} kr/month</div>
          </div>
          <div class="info-pair">
            <div class="info-pair-key">Consumption Price</div>
            <div class="info-pair-val">
              {{
                property.consumption_price !== null
                  ? consumption_price
                  : "Unspecified"
              }}
              {{ property.consumption_price !== null ? "kr/year" : "" }}
            </div>
          </div>
          <div class="info-pair">
            <div class="info-pair-key">Price/m²</div>
            <div class="info-pair-val">{{ property.price_sqm }} kr/m²</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["property"],
};
</script>

<style scoped>
h2 {
  padding: 0;
  margin: 0;
  font-size: 25px;
}
a {
  color: #014124;
  font-weight: bold;
}
a:visited {
  text-decoration: none;
  color: #014124;
}
h3 {
  padding: 0;
  margin: 0;
  margin-top: 5px;
  margin-bottom: 5px;
  font-size: 30px;
  font-weight: 100;
}
small {
  padding: 0;
  margin: 0;
  font-size: 14px;
  color: #575656;
}
img {
  max-width: 100%;
}
p {
  margin: 0;
  padding: 0;
}

.container {
  display: flex;
  flex-direction: column;
}

.info-container {
  margin-right: 10px;
  margin-left: 10px;
}

.info-street {
  margin-bottom: 10px;
  border-bottom: 1px solid #e2e2e2;
}
.info-link {
    margin-bottom: 20px;
}
.info-property {
  border-bottom: 1px solid #e2e2e2;
  margin-bottom: 10px;
}
.info-pair {
  display: flex;
  justify-content: space-between;
  margin-bottom: 5px;
}
.info-pair:last-child {
  margin-bottom: 15px;
}
.ino-par-key {
  width: 50%;
  font-weight: lighter;
}
.info-pair-val {
  width: 50%;
  justify-self: flex-start;
  font-weight: bold;
}
</style>