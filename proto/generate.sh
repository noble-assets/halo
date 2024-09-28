cd proto
buf generate --template buf.gen.gogo.yaml
buf generate --template buf.gen.pulsar.yaml
cd ..

cp -r github.com/noble-assets/halo/v2/* ./
cp -r api/halo/* api/
find api/ -type f -name "*.go" -exec sed -i 's|github.com/noble-assets/halo/v2/api/halo|github.com/noble-assets/halo/v2/api|g' {} +

rm -rf github.com
rm -rf api/halo
rm -rf halo
