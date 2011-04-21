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
        source='''\
        pkg/hdf5.go
        pkg/h5d.go
        pkg/h5f.go
        pkg/h5g.go
        pkg/h5p.go
        pkg/h5s.go
        pkg/h5t.go
        ''',
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

    ctx(
        features='go goprogram',
        name   = 'test-go-cpxcmpd',
        source ='cmd/test-go-cpxcmpd.go',
        target = 'test-go-cpxcmpd',
        use = ['go-hdf5',],
        )
