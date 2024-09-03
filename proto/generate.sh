cd proto
buf generate --template buf.gen.gogo.yaml
buf generate --template buf.gen.pulsar.yaml
cd ..

cp -r github.com/noble-assets/halo/* ./
rm -rf github.com
cp -r api/halo/* api/

find api/ -type f -name "*.go" -exec sed -i 's|github.com/noble-assets/halo/api/halo|github.com/noble-assets/halo/api|g' {} +

rm -rf api/halo
rm -rf halo
