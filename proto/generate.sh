cd proto
buf generate --template buf.gen.gogo.yaml
cd ..

cp -r github.com/noble-assets/halo/* ./
rm -rf github.com
