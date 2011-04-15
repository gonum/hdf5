#

top = '.'
out = '__build__'

def options(ctx):
    pass

def configure(ctx):
    ctx.load('go')
    
def build(ctx):

    ctx(
        features='cgopackage',
        name ='go-hdf5',
        source='pkg/hdf5.go',
        target='hdf5',
        use = [
            'hdf5',
            ],
        )

    ctx(
        features='go goprogram',
        name   = 'test-go-hdf5',
        source ='cmd/test-go-hdf5.go',
        target = 'test-go-hdf5',
        use = ['go-hdf5',],
        )
