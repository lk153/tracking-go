from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task(3)
    def fetch_product_pb(self):
        self.client.get("/v1/products?limit=10")

    @task(3)
    def fetch_product(self):
        self.client.get("/products?limit=10")

    @task(1)
    def get_home(self):
        self.client.get("/")

    @task(1)
    def get_home(self):
        self.client.get("/v1/product/1")
        for item_id in range(1, 10):
            self.client.get(f"/v1/product/{item_id}", name="/item")
            time.sleep(1)

    @task(1)
    def create_product(self):
        response = self.client.post("/v1/product", json={"data": {"name": "Huawei 1", "price": 3000, "type": "simple", "status": 1}})
        print("Response status code:", response.status_code)
        print("Response text:", response.text)      